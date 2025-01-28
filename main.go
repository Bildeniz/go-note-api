package main

import (
	"Go-Note-API/Endpoints"
	"Go-Note-API/Models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := Models.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	Models.MakeMigrations(db)

	router := gin.Default()

	Endpoints.RegisterEndpoints(router)

	router.Run(":8080")
}
