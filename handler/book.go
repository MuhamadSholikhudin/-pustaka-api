package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erorrs": err,
		})
		return
	}
	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := converttoBookResponse(b)
		// bookResponse := book.BookResponse{
		// 	ID:          b.ID,
		// 	Title:       b.Title,
		// 	Description: b.Description,
		// 	Rating:      b.Rating,
		// 	Discount:    b.Discount,
		// }
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erorrs": err,
		})
		return
	}

	//Convert
	bookResponse := converttoBookResponse(b)
	// bookResponse := book.BookResponse{
	// 	ID:          b.ID,
	// 	Title:       b.Title,
	// 	Description: b.Description,
	// 	Rating:      b.Rating,
	// 	Discount:    b.Discount,
	// }

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

// func PostbookHandler(c *gin.Context) {
func (h *bookHandler) CreateBook(c *gin.Context) {
	///title, price

	// var bookInpput book.BookInput
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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
	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"title": bookRequest.Title,
	// 	"price": bookRequest.Price,
	// })
}

// func PostbookHandler(c *gin.Context) {
func (h *bookHandler) UpdateBook(c *gin.Context) {
	///title, price

	// var bookInpput book.BookInput
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"title": bookRequest.Title,
	// 	"price": bookRequest.Price,
	// })
}

func converttoBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
