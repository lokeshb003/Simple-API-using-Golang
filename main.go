package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"errors"
)

type book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

var books = []book{
	{ID: "1", Title: "Phoenix Project", Name: "Lokesh Balaji"},
	{ID: "2", Title: "Unicorn Project", Name: "Lawanya Balaji"},
}

func getallbooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func main() {
	router := gin.Default()
	router.GET("/books", getallbooks)
	router.Run("localhost:8080")
}
