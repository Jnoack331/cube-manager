package main

import (
	"cube-manager/src/controllers"
	"cube-manager/src/minecraft-manager"
	"embed"
	_ "embed"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/* assets/js/* assets/img/*
var embedFS embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	minecraftServer := minecraft_manager.NewMinecraftManager("server.jar")
	if os.Getenv("CUBE_AUTORUN_SERVER") == "true" {
		minecraftServer.Start()
	}

	r := gin.Default()
	templ := template.Must(template.New("").ParseFS(embedFS, "templates/*.tmpl"))
	r.SetHTMLTemplate(templ)
	r.StaticFS("/public", http.FS(embedFS))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 0,
		Path:   "/",
	})
	r.Use(sessions.Sessions("sessions", store))

	r.GET("/", controllers.Dashboard)
	r.GET("/filelist", controllers.Filelist)
	r.POST("/upload", controllers.Upload)
	r.POST("/delete", controllers.Delete)
	r.GET("/download", controllers.Download)
	r.POST("/logout", controllers.Logout)
	r.POST("/restart", controllers.Restart)
	r.GET("/authenticated", controllers.Authenticated)

	r.GET("/server/output", func(context *gin.Context) {
		controllers.ServerOutput(context, minecraftServer)
	})
	r.POST("/server/command", func(context *gin.Context) {
		controllers.ServerCommand(context, minecraftServer)
	})
	r.GET("/server/status", func(c *gin.Context) {
		controllers.ServerStatus(c, minecraftServer)
	})

	r.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	r.POST("/login", controllers.Login)

	r.Run()
}
