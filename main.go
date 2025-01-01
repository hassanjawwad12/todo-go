package main

import (
	"errors"
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

func GetbyId(id string) (*Todo, error) {

	// Iterate over the todo to find the one with the requested id
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("no todo exist against this id")
}

func GetTodo(context *gin.Context) {

	// Get the id from the param
	id := context.Param("id")

	todo, err := GetbyId(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func DeleteTodo(context *gin.Context) {

	id := context.Param("id")

	index := -1
	for i, todo := range todos {
		if todo.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Remove the todo from the array
	todos = append(todos[:index], todos[index+1:]...)

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
func main() {

	fmt.Println("Welcome to TODO APP")
	router := gin.Default()
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodo)
	router.POST("/todos", AddTodo)
	router.DELETE("/todos/:id", DeleteTodo)
	router.Run("localhost:8080") // Run our server

}
