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
	router.LoadHTMLGlob("static/*")
	router.Static("/images", "./images")
	router.Static("/audio", "./audio")
	router.Static("/bookImg", "./bookImg")
	router.Static("/static", "./static")
	router.Static("/css", "./css")
	router.GET("/", controllers.HelloWorld)
	router.GET("/add", controllers.CreateBook)
	router.GET("/api/get/book/list/", controllers.GetBookList)
	router.GET("/api/get/book/detail/:id", controllers.GetBook)
	router.POST("/api/create/book/", controllers.CreateBook)
	router.POST("/api/delete/book/:id", controllers.DeleteBook)

	router.Run(":4444")
}
