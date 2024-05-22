package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []todo{
	{ID: "1", Title: "Learn Go", Status: "In Progress"},
	{ID: "2", Title: "Read a book", Status: "Pending"},
	{ID: "3", Title: "Walk the dog", Status: "Done"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func addTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodoStatus(c *gin.Context) {
	id := c.Param("id")
	var updatedStatus struct {
		Status string `json:"status"`
	}

	if err := c.BindJSON(&updatedStatus); err != nil {
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos[i].Status = updatedStatus.Status
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func deleteTodoByID(c *gin.Context) {
	id := c.Param("id")

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", addTodo)
	router.PUT("/todos/:id", updateTodoStatus)
	router.DELETE("/todos/:id", deleteTodoByID)
	router.Run("localhost:8080")
}
