package book

import "time"

type Book struct {
	ID          int
	Title       string
	Descrition  string
	Price       int
	Rating      int
	Discount    int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
