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

	// bookFilrepository := book.NewFileRepository()

	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	//Versioning API
	v1 := gin.Default()

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	v1.Run()

}
