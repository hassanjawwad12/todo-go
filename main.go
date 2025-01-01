package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id   string `json: "id", required`
	Note string `json: "note", required`
}

// Created array of 3 todos
var todos = []Todo{
	{Id: "1", Note: "Complete the project"},
	{Id: "2", Note: "Make the pull request"},
	{Id: "3", Note: "Read a book by Dostoevsky"},
}

func GetTodos(context *gin.Context) {

}

func main() {

	router := gin.Default()
	router.GET("/todos")
	router.Run("localhost:8080") // Run our server

	// The server only accepts JSON

}
