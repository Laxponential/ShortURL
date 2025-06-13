package main

import (
	"shortURL/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// create router
	r := gin.Default()

	// two handlers
	// post url to shortenize the url
	// get url to fetch the url

	r.POST("/", handler.ShortenURL())
	r.GET("/")

	r.Run(":8123")
}
