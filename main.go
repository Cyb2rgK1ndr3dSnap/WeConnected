package main

import (
	"net/http"

	"github.com/Cyb2rgK1ndr3dSnap/we-connected/database"
	"github.com/Cyb2rgK1ndr3dSnap/we-connected/helper"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.ConnectToDB()
	if err != nil {
		panic("Error al conectar a la base de datos")
	}
	defer database.CloseDB(db) // Cierra la conexión al final de la aplicación

	// Create a new Gin router
	routes := gin.Default()
	// Define a route for the "ping" endpoint
	routes.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server := &http.Server{
		Addr:    ":8000",
		Handler: routes,
	}
	// Run the server on port 8000
	//routes.Run(":8000")
	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
