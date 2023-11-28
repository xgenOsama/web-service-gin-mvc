package main

import (
	"example.com/gin-mvc/controllers"
	"example.com/gin-mvc/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostByID)
	r.POST("/posts/:id", controllers.PostsUpdate)
	r.POST("/posts/delete/:id", controllers.PostDELETE)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
