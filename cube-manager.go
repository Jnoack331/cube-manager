package main

import (
	"cube-manager/src/controllers"
	"embed"
	_ "embed"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/* assets/js/* assets/img/*
var embedFS embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	r.POST("/", controllers.Dashboard)
	r.GET("/filelist", controllers.Filelist)
	r.POST("/upload", controllers.Upload)
	r.POST("/delete", controllers.Delete)
	r.GET("/download", controllers.Download)
	r.POST("/logout", controllers.Logout)
	r.POST("/restart", controllers.Restart)
	r.GET("/authenticated", controllers.Authenticated)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.POST("/login", controllers.Login)
	r.Run()
}
