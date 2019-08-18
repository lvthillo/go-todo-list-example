package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvthillo/go-todo-list-example/handlers"
)

func main() {
	// create Gin server.
	// it returns an object that we can use to configure and run the web server.
	r := gin.Default()
	//GET: This endpoint enables users to retrieve the entire to-do list.
	r.GET("/todo", handlers.GetTodoListHandler)
	//POST: This endpoint enables users to add new items to the list.
	r.POST("/todo", handlers.AddTodoHandler)
	//DELETE: This endpoint enables users to delete a to-do from the list based on an ID.
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	//PUT: This endpoint enables users to change a to-do item from incomplete to complete.
	r.PUT("/todo", handlers.CompleteTodoHandler)
	//GET: This endpoint enables users to retrieve one item of the todo-list based on ID.
	r.GET("/todo/:id", handlers.GetTodoItemHandler)

	// run server on 3000 and panics when error occurs.
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
