package main

import "github.com/gin-gonic/gin"

func main() {

	// create router
	r := gin.Default()

	// two handlers
	// post url to shortenize the url
	// get url to fetch the url

	r.POST("/")
	r.GET("/")

	r.Run(":8123")
}
