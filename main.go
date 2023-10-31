package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	//"encoding/json"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "I Found You", Author: "stephen nicolas", Quantity: 200},
	{ID: "2", Title: "mercy and love", Author: "marviigrey castle", Quantity: 1000},
	{ID: "3", Title: "I Found You", Author: "Joveka grant", Quantity: 200},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) //an HTTP response writer by the gin package.
	//prints out result of the89*
}

func createBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks) //Get requests of /books will run the function getBooks
	router.POST("/books", createBooks)
	router.GET("books/:id", bookById)
	router.Run("localhost:8085") //port and ip to run on
}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
