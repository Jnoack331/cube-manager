package main

import (
	Frontend "cube-manager/frontend"
	"cube-manager/src/controllers"
	"cube-manager/src/minecraft-manager"
	_ "embed"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request, minecraftServer *minecraft_manager.MinecraftManager) {
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for true {
		output := minecraftServer.GetOutput()
		conn.WriteMessage(1, []byte(output))
		time.Sleep(200 * time.Millisecond)
	}
}

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
	templ := template.Must(template.New("").ParseFS(Frontend.EmbedFS, "dist/index.html"))
	r.SetHTMLTemplate(templ)
	assets, _ := fs.Sub(Frontend.EmbedFS, "dist/assets")
	r.StaticFS("/assets", http.FS(assets))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 0,
		Path:   "/",
	})
	r.Use(sessions.Sessions("sessions", store))

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Headers", "content-type")
		c.Header("Access-Control-Allow-Credentials", "true")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/ws/logs", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("authenticated") != true {
			c.Redirect(302, "/login")
		}
		wshandler(c.Writer, c.Request, minecraftServer)
	})

	r.GET("/filelist", controllers.Filelist)
	r.POST("/upload", controllers.Upload)
	r.POST("/delete", controllers.Delete)
	r.GET("/download", controllers.Download)
	r.POST("/logout", controllers.Logout)
	r.POST("/restart", func(c *gin.Context) {
		controllers.Restart(c, minecraftServer)
	})
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

	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Data(200, gin.MIMEPlain, nil)
			return
		}
		c.Data(404, gin.MIMEPlain, nil)
	})

	r.Run()
}
