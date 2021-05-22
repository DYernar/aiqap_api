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
	books, _ := db.GetAllBooks()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
		"books": books,
	})
}

var categories = []string{
	"Драма",
	"Детектив",
	"Қиял ғажайып",
	"Тарихи",
	"Ғылыми",
	"Поэзия",
	"Балаларға арналған",
	"Діни",
	"Шет елдің кітаптары",
	"Психология",
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
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "addbook.tmpl", gin.H{
			"categories": categories,
		})
	} else if c.Request.Method == "POST" {
		c.Request.Header.Add("content-type", "application/json")

		file, header, err := c.Request.FormFile("audio")
		if err != nil {
			fmt.Println(err)
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
		c.Request.ParseForm()
		book.AudioLink = "/audio/" + filename + ".mp3"
		book.Title = c.Request.FormValue("title")
		book.AuthorName = c.Request.FormValue("author_name")
		book.AuthorSurname = c.Request.FormValue("author_surname")
		book.Description = c.Request.FormValue("description")
		book.Categories = c.Request.Form["categories"]

		json.NewDecoder(c.Request.Body).Decode(&book)

		result := db.CreateBook(book)
		json.NewEncoder(c.Writer).Encode(result)
	}

}

func getID(objectID string) string {
	l := strings.Split(objectID, "\"")
	return l[1]
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
	c.Request.Header.Add("content-type", "application/json")
	var book models.Book

	err := c.ShouldBindUri(&book)
	if err != nil {
		c.String(http.StatusNotFound, "book not found "+err.Error())
		return
	}
	fmt.Println(book.ID)
	docID, err := primitive.ObjectIDFromHex(book.ID)
	res := db.DeleteBook(docID)
	if res {
		c.String(http.StatusOK, "deleted book successfully")
	} else {
		c.String(http.StatusOK, "book is not deleted")
	}
}

func UpdateBook(c *gin.Context) {
	c.String(http.StatusOK, "get book page")
}
