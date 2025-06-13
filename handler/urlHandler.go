package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const urlSize int = 6

func generateShortURL(n int) string {

	b := make([]byte, n)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:n]

	// rand.Read() + base64 : Safer entropy distribution, Avoids bias from poor random selection logic
}

func ShortenURL(rdb *redis.Client) gin.HandlerFunc {
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

		slug := generateShortURL(urlSize)

		shortURL := fmt.Sprintf("short.ly/%s", slug)

		ctx := context.Background()
		// store to redis
		if err := rdb.Set(ctx, shortURL, urlBody.Url, 0).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not save to redis",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "URL generated: " + shortURL,
		})
	}
}
