package main

import (
	"os"
	"shortURL/cache"
	"shortURL/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load() // only needed for local dev (non-docker runs)
}

func main() {

	// create router
	r := gin.Default()

	// Init Redis
	redisClient := cache.InitRedis()

	// two handlers
	// post url to shortenize the url
	// get url to fetch the url
	r.POST("/", handler.ShortenURL(redisClient))
	r.GET("/:slug", handler.ResolveURL(redisClient))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8123"
	}

	r.Run(":" + port)
}
