package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//routing default
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Muhammad Sholikhudin",
			"bio":  "I am Seoarang prormmer",
		})
	}) // route url
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"content":  "Hello",
			"su_title": "World",
		})
	}) // route url

	router.Run()

}
