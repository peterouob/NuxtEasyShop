package service

import (
	"easyshop/database"
	"easyshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	u := model.User{}
	err := c.ShouldBind(&u)
	if err != nil {
		return
	}
	result := database.DB.Where("username = ? and password = ?", u.Username, u.Password).Find(&u)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"token": "token",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "not found user",
		})
		return
	}
}
