package service

import (
	"easyshop/database"
	"easyshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	u := model.User{}
	c.ShouldBind(&u)
	res := database.DB.Where("username = ? or email = ?", u.Username, u.Email).Find(&u)
	if res.RowsAffected >= 1 {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "have user",
		})
		return
	} else {
		database.DB.Create(&u)
		c.JSON(http.StatusOK, gin.H{
			"token": "123",
			"msg":   "ok",
		})
	}
}
