package Endpoints

import (
	"Go-Note-API/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type notes struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

func connectOr500(c *gin.Context) *gorm.DB {
	db, err := Models.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection failed"})
		c.Abort()
		return nil
	}
	return db
}

func RegisterEndpoints(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/notes", func(c *gin.Context) {
			db := connectOr500(c)

			var notes []notes
			results := db.Find(&notes)

			if results.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No Notes founded"})
			}

			c.JSON(http.StatusOK, notes)
		})
	}
}
