package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"authenticated": false,
		})
		return
	}

	if form.Username != os.Getenv("CUBE_USERNAME") || form.Password != os.Getenv("CUBE_PASSWORD") {
		c.JSON(http.StatusOK, gin.H{
			"authenticated": false,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("authenticated", true)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
	})
}

func Authenticated(c *gin.Context) {
	session := sessions.Default(c)
	authenticated := false
	if session.Get("authenticated") == true {
		authenticated = true
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": authenticated,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	session.Delete("authenticated")
	session.Save()
	c.JSON(http.StatusOK, gin.H{})
}
