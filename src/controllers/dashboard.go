package controllers

import (
	Filesystem "cube-manager/src/filesystem"
	MinecraftManager "cube-manager/src/minecraft-manager"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	minecraft := MinecraftManager.MinecraftManager{}
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	currentDirectory, _ := os.Getwd()
	value := c.Query("path")
	currentDirectoryLength := len(strings.Split(currentDirectory, "/"))
	pathDirectoryLength := len(strings.Split(value, "/"))
	if len(value) > 0 && currentDirectoryLength < pathDirectoryLength {
		currentDirectory = value
	}

	filesystem := Filesystem.Filesystem{}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"running":    minecraft.IsServerRunning(),
		"filelist":   filesystem.GetFileList(currentDirectory),
		"currenPath": value,
	})
}

func Upload(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
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

	currentPath := c.Query("currentPath")
	destinationFile := currentPath + "/" + filename
	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}

	c.Redirect(302, "/?path="+currentPath)

	if err != nil {
		log.Fatal(err)
	}
}

func Delete(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	currentPath := c.Query("currentPath")
	path := c.Query("path")
	os.RemoveAll(path)
	c.Redirect(302, "/?path="+currentPath)
}

func Download(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	currentPath := c.Query("currentPath")
	path := c.Query("path")
	splitted := strings.Split(path, "/")
	fileName := splitted[len(splitted)-1]

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.File(path)
	c.Redirect(302, "/?path="+currentPath)
}

func Restart(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	cmd := exec.Command("sudo", "systemctl", "restart", "minecraft")
	cmd.Run()
	currentPath := c.Query("currentPath")
	c.Redirect(302, "/?path="+currentPath)
}
