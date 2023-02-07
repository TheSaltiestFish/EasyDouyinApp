package main

import (
	"github.com/TheSaltiestFish/EasyDouyinApp/dal"
	"github.com/TheSaltiestFish/EasyDouyinApp/service"
	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)
	dal.Init()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
