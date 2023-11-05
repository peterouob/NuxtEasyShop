package router

import (
	"easyshop/helper"
	"easyshop/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.Use(helper.Cors)
	r.POST("/user/login", service.Login)
	r.POST("/user/register", service.Register)
}
