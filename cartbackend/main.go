package main

import (
	"easyshop/database"
	"easyshop/router"
	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		database.MysqlInit()
	}()
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":90")
}
