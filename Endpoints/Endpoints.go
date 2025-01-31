package Endpoints

import (
	"Go-Note-API/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

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
		api.GET("/notes", func(c *gin.Context) { // Get all notes
			db := connectOr500(c)

			var notes []Models.Notes
			results := db.Find(&notes)

			if results.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No Notes founded"})
				return
			}

			c.JSON(http.StatusOK, notes)
		})

		api.POST("/notes", func(c *gin.Context) {
			var newNote Models.Notes

			if err := c.ShouldBindJSON(&newNote); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			db := connectOr500(c)

			result := db.Create(&newNote)

			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
				return
			}

			c.JSON(http.StatusCreated, newNote)
		})

		api.GET("/notes/:id", func(c *gin.Context) {
			id := c.Param("id")
			db := connectOr500(c)

			var note Models.Notes
			results := db.First(&note, id)
			if results.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
				return
			}

			c.JSON(http.StatusOK, note)
		})

		api.DELETE("/notes/:id", func(c *gin.Context) {
			id := c.Param("id")
			db := connectOr500(c)

			result := db.Delete(&Models.Notes{}, id)

			if result.RowsAffected == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
				return
			} else if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
		})

		api.PUT("/notes/:id", func(c *gin.Context) {
			id := c.Param("id")
			db := connectOr500(c)

			var note Models.Notes
			result := db.First(&note, id)

			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Not Founded"})
				return
			}

			if err := c.ShouldBindJSON(&note); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
				return
			}

			db.Save(&note)
			c.JSON(http.StatusOK, note)
		})
	}
}
