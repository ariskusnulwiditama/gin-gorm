package main

import (
	"github.com/ariskusnulwiditama/gin-gorm/controllers"
	"github.com/ariskusnulwiditama/gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDataBase()
	route.GET("/books", controllers.FindBooks)
	route.GET("/books/:id", controllers.FindBook)
	route.POST("/books", controllers.CreateBook)
	route.PATCH("/books/:id", controllers.UpdateBook)
	route.DELETE("/books/:id", controllers.DeleteBook)
	route.Run()
}
