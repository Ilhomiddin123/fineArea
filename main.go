package main

import (
	"fineArea/controllers"
	"fineArea/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.POST("/upload", controllers.UploadVehicle)

	err := r.Run(":8888")
	fmt.Println("Server is running on port 8888")
	if err != nil {
		panic("Error while starting server")
	}
}
