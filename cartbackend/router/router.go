package router

import (
	"easyshop/helper"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetupRouter(r *gin.Engine) {
	r.Use(helper.Cors)
	r.POST("/user/login", login)
}

func login(c *gin.Context) {
	u := User{}
	err := c.ShouldBind(&u)
	if err != nil {
		return
	}
	if u.Username == "123" && u.Password == "123" {
		c.JSON(200, gin.H{
			"token": "123",
			"u":     u,
		})
	} else {
		c.JSON(401, "not auth")
	}
}
