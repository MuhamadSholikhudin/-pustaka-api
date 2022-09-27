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

	// Auto migration
	// db.AutoMigrate(&book.Book{})

	//CRUD

	///Create
	// book := book.Book{}
	// book.Title = "Man Tiger"
	// book.Price = 90000
	// book.Rating = 101
	// book.Description = "Ini adalah Film Man Tiger"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Creating book")
	// }

	// Read

	// var book book.Book

	// // err = db.Debug().First(&book).Error // Get By Last
	// err = db.Debug().First(&book, 1).Error // Get By id
	// if err != nil {
	// 	fmt.Println("Error Read First book")
	// }

	// fmt.Println("Output Title : ", book.Title)
	// fmt.Println("Output Price : ", book.Price)
	// fmt.Println("Output Rating : ", book.Rating)
	// fmt.Println("Output Description", book.Description)

	var books []book.Book

	// err = db.Debug().Find(&books, 1).Error // GET data ROWS
	err = db.Debug().Where("title = ?", "Man Tiger").Find(&books).Error // GET data ROWS WHERE
	if err != nil {
		fmt.Println("Error Read First book")
	}

	for _, b := range books {
		fmt.Println("Output Book : ", b.Title)

	}

	//Versioning API
	v1 := gin.Default()
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostbookHandler)

	v1.Run()

}
