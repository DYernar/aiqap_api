package controllers

import (
	"aiqap-back/db"
	"aiqap-back/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HelloWorld(c *gin.Context) {
	c.String(200, "Hello World")
}

func GetBookList(c *gin.Context) {
	books, err := db.GetAllBooks()
	if err != nil {
		fmt.Println(err)
		c.String(500, err.Error())
		return
	}
	json.NewEncoder(c.Writer).Encode(books)

}

func CreateBook(c *gin.Context) {
	c.Request.Header.Add("content-type", "application/json")

	file, header, err := c.Request.FormFile("audio")
	if err != nil {
		c.String(400, "no file provided")
		return
	}

	filename := header.Filename
	filename = strings.ReplaceAll(filename, " ", "")
	out, err := os.Create("./audio/" + filename + ".mp3")
	if err != nil {
		fmt.Println(err)
		c.String(500, "internal server error")
		return
	}
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
		c.String(500, "internal server error")
		return
	}
	defer out.Close()

	var book models.Book
	book.AudioLink = "/audio/" + filename + ".mp3"
	book.Title = c.Request.FormValue("title")
	book.AuthorName = c.Request.FormValue("author_name")
	book.AuthorSurname = c.Request.FormValue("author_surname")
	book.Description = c.Request.FormValue("description")
	book.Categories = strings.Split(c.Request.FormValue("categories"), ", ")

	json.NewDecoder(c.Request.Body).Decode(&book)

	result := db.CreateBook(book)
	json.NewEncoder(c.Writer).Encode(result)
}

func GetBook(c *gin.Context) {
	c.Request.Header.Add("content-type", "application/json")
	rawId := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(rawId)

	book, err := db.GetBook(id)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	json.NewEncoder(c.Writer).Encode(book)

}

func DeleteBook(c *gin.Context) {
	c.String(http.StatusOK, "delete book page")
}

func UpdateBook(c *gin.Context) {
	c.String(http.StatusOK, "get book page")
}
