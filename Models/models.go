package Models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Notes struct {
	gorm.Model
	Title   string
	Content string
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
		return nil, err
	}

	err = db.AutoMigrate(&Notes{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return nil, err
	}

	return db, nil
}
