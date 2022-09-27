package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connectin Error")

	}
	fmt.Println("Database Connected")
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	//Versioning API
	v1 := gin.Default()
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostbookHandler)

	v1.Run()

}
