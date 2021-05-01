package main

import (
	"aiqap-back/controllers"

	"github.com/gin-gonic/gin"
)

// to connect to aws shell:
//

func main() {
	router := gin.Default()
	router.Static("/audio", "./audio")
	router.GET("/api/get/book/list/", controllers.GetBookList)
	router.GET("/api/get/book/detail/:id", controllers.GetBook)
	router.POST("/api/create/book/", controllers.CreateBook)
	router.DELETE("/api/delete/book/:id", controllers.DeleteBook)

	router.Run()
}
