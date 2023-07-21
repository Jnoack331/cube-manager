package controllers

import (
	"cube-manager/src/filesystem"
	minecraft_manager "cube-manager/src/minecraft-manager"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func Filelist(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"filelist": []string{},
		})
		return
	}

	myfilesystem := filesystem.Filesystem{}
	currentDirectory, _ := os.Getwd()
	requestedPath := c.Query("path")

	if myfilesystem.IsInRootPath(requestedPath) {
		currentDirectory = requestedPath
	}

	c.JSON(http.StatusOK, gin.H{
		"filelist":    myfilesystem.GetFileList(currentDirectory),
		"currentPath": currentDirectory,
	})
}

func Upload(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
		return
	}

	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	fmt.Println(header.Filename)
	tmpPath := "/tmp/" + filename
	out, err := os.Create(tmpPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)

	input, err := ioutil.ReadFile(tmpPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetPath := c.Query("path")
	destinationFile := targetPath + "/" + filename
	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	if err != nil {
		log.Fatal(err)
	}
}

type DeleteForm struct {
	Path string `json:"path" binding:"required"`
}

func Delete(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
		return
	}

	var form DeleteForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{})
	}

	path := form.Path
	myFilesystem := filesystem.Filesystem{}
	myFilesystem.Delete(path)
	c.JSON(http.StatusOK, gin.H{})
}

func Download(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
		return
	}

	path := c.Query("path")
	splitted := strings.Split(path, "/")
	fileName := splitted[len(splitted)-1]

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.File(path)
}

func Restart(c *gin.Context, server *minecraft_manager.MinecraftManager) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
		return
	}

	if server.IsServerRunning() {
		server.Stop()
		for server.IsServerRunning() {
			time.Sleep(500 * time.Millisecond)
		}
	}

	server.Start()
	c.JSON(http.StatusOK, gin.H{})
}
