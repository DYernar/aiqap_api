package main

import (
	"aiqap-back/controllers"
	"aiqap-back/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

// to connect to aws shell:
//

func main() {
	_, err := db.GetAllBooks()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	router.Static("/audio", "./audio")
	router.GET("/", controllers.HelloWorld)
	router.GET("/api/get/book/list/", controllers.GetBookList)
	router.GET("/api/get/book/detail/:id", controllers.GetBook)
	router.POST("/api/create/book/", controllers.CreateBook)
	router.DELETE("/api/delete/book/:id", controllers.DeleteBook)

	router.Run()
}
