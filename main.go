package main

import (
	"Go-Note-API/Models"
	"log"
)

func main() {
	db, err := Models.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&Models.Notes{Title: "New Note", Content: "Hello World"})
}
