package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		type UrlBody struct {
			Url string `json:"url"`
		}

		var urlBody UrlBody
		if err := c.BindJSON(&urlBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid URL",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "URL Succesfully received",
		})
	}
}
