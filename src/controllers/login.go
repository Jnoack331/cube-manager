package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"error": "Wrong Credentials!",
		})
		return
	}

	if form.User != os.Getenv("USERNAME") || form.Password != os.Getenv("PASSWORD") {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"error": "Wrong Credentials!",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("authenticated", true)
	session.Save()
	c.Redirect(302, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("authenticated") != true {
		c.Redirect(302, "/login")
	}

	session.Delete("authenticated")
	session.Save()
	c.Redirect(302, "/login")
}
