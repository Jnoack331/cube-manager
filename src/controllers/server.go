package controllers

import (
	minecraft_manager "cube-manager/src/minecraft-manager"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerOutput(c *gin.Context, minecraftServer *minecraft_manager.MinecraftManager) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	c.JSON(http.StatusOK, gin.H{
		"output": minecraftServer.GetOutput(),
	})
}

type CommandForm struct {
	Command string `json:"command" binding:"required"`
}

func ServerCommand(c *gin.Context, minecraftServer *minecraft_manager.MinecraftManager) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	var form CommandForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{})
	}
	minecraftServer.SendCommand(form.Command)

	c.JSON(http.StatusOK, gin.H{})
}

func ServerStatus(c *gin.Context, minecraftServer *minecraft_manager.MinecraftManager) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	c.JSON(http.StatusOK, gin.H{
		"running": minecraftServer.IsServerRunning(),
	})
}
