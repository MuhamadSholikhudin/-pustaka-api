package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Muhammad Sholikhudin",
		"bio":  "I am Seoarang prormmer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content":  "Hello",
		"su_title": "World",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostbookHandler(c *gin.Context) {
	///title, price

	var bookInpput book.BookInput

	err := c.ShouldBindJSON(&bookInpput)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition :%s", e.Field(), e.ActualTag())
			// c.JSON(http.StatusBadRequest, errorMessage)
			// return
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		// c.JSON(http.StatusBadRequest, err)
		// fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":     bookInpput.Title,
		"price":     bookInpput.Price,
		"sub_title": bookInpput.SubTitle,
	})
}
