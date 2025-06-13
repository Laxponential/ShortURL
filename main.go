package main

import (
	"shortURL/cache"
	"shortURL/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// create router
	r := gin.Default()

	// Init Redis
	redisClient := cache.InitRedis()

	// two handlers
	// post url to shortenize the url
	// get url to fetch the url
	r.POST("/", handler.ShortenURL(redisClient))
	r.GET("/")

	r.Run(":8123")
}
