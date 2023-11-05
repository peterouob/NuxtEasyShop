package main

import (
	"easyshop/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":90")
}
