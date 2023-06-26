package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Description string `json:"description"`
	Author string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}

var books = []book{
	{ID: "1", Title: "Cutting for Stone: A Novel", Description: "Cutting for Stone: A Novel", Author: " Abraham Verghese", Price: 8.37},
	{ID: "2", Title: "How to Write Clearly", Description: "Write with purpose, reach your reader and make your meaning crystal clear", Price: 16.00},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range books {
			if a.ID == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
			return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}