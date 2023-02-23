package main

import (
	"net/http"

	"github.com/devesh/gin-gorm-crud/controllers"
	"github.com/devesh/gin-gorm-crud/helper"
	"github.com/devesh/gin-gorm-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("started Server!") //---log history
	routes := gin.Default()

	models.ConnectDatabase() // calling the database connection

	routes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello i am ready ",
		})
	})

	routes.GET("/books", controllers.FindBooks)
	routes.POST("/books", controllers.CreateBook)
	routes.GET("/books/:id", controllers.FindBook)
	routes.PUT("/books/:id", controllers.UpdateBook)
	routes.DELETE("/books/:id", controllers.DeleteBook)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
