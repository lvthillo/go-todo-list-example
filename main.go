package main

import (
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/lvthillo/go-todo-list-example/handlers"
)

func main() {
	// create Gin server.
	// it returns an object that we can use to configure and run the web server.
	r := gin.Default()
	// routing in Gin is specific.
	// if you have a configuration like /* because this will interfere with every other route in your web server.
	// this is possible in NodeJS where the determination is based on the most specific till least specific paths.
	//  /api/something would have precedence over /*.
	// To 'hack' this in Gin we use the NoRoute function.
	// It matches all routes that have not been specified already. This route function will assume that this call is asking for a file and attempt to find this file.
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		// this route function will assume that this call is asking for a file and attempt to find this file.
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			// if not found, serve index.html
			c.File("./ui/dist/ui/index.html")
		} else {
			// if it's found, serve file.
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	// all endpoints are strucutred in the same manner r.<METHOD>(<PATH>, <Gin function>)
	// the Gin function is basically any function that takes the parameter of a gin.Context pointer (can contains info, cookies, etc)
	// if you look at the NoRoute function, you will see an example of an anonymous function with the input of a gin.Context pointer.
	// check https://godoc.org/github.com/gin-gonic/gin

	//GET: This endpoint enables users to retrieve the entire to-do list.
	r.GET("/todo", handlers.GetTodoListHandler)
	//POST: This endpoint enables users to add new items to the list.
	r.POST("/todo", handlers.AddTodoHandler)
	//DELETE: This endpoint enables users to delete a to-do from the list based on an ID.
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	//PUT: This endpoint enables users to change a to-do item from incomplete to complete.
	r.PUT("/todo", handlers.CompleteTodoHandler)

	// run server on 3000 and panics when error occurs.
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
