package main

import (
	"example.com/gin-mvc/initializers"
	"example.com/gin-mvc/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
