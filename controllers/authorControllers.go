package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthor(c *gin.Context) {
	c.String(http.StatusOK, "get book page")
}
