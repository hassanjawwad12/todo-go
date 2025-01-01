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

// Context contains the information about the http request
func GetTodos(context *gin.Context) {

	// The server only accepts JSON so we  Transform the data into the JSON
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(context *gin.Context) {

	var newTodo Todo

	// Bind json in the request body to the newTodo
	err := context.BindJSON(&newTodo)
	if err != nil {
		fmt.Println("Error has occured while adding new todo", err)
		return
	}

	// Added the new todo in the array
	todos = append(todos, newTodo)

	// Comvert into JSON
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {

	fmt.Println("Welcome to TODO APP")
	router := gin.Default()
	router.GET("/todos", GetTodos)
	router.POST("/todos", AddTodo)
	router.Run("localhost:8080") // Run our server

}
