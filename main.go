package main

import (
	"fmt"
	"log"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// //routing default
	// router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name": "Muhammad Sholikhudin",
	// 		"bio":  "I am Seoarang prormmer",
	// 	})
	// }) // route url
	// router.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"content":  "Hello",
	// 		"su_title": "World",
	// 	})
	// }) // route url

	// router.GET("/", rootHandler)
	// router.GET("/hello", helloHandler)
	// router.GET("/books/:id/:title", bookHandler)
	// router.GET("/query", queryHandler)
	// router.POST("/books", postbookHandler)

	//Connection DB MYSQL
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connectin Error")

	}
	fmt.Println("Database Connected")

	//Versioning API
	v1 := gin.Default()
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostbookHandler)

	v1.Run()

}
