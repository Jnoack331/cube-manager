package controllers

import (
	"cube-manager/src/filesystem"
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
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func Filelist(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"filelist": []string{},
		})
	}

	myfilesystem := filesystem.Filesystem{}
	currentDirectory, _ := os.Getwd()
	value := c.Query("path")
	currentDirectoryLength := len(strings.Split(currentDirectory, "/"))
	pathDirectoryLength := len(strings.Split(value, "/"))
	if len(value) > 0 && currentDirectoryLength < pathDirectoryLength {
		currentDirectory = value
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

func Delete(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	c.Request.ParseForm()
	path := c.Request.Form.Get("path")
	os.RemoveAll(path)
	c.JSON(http.StatusOK, gin.H{})
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
