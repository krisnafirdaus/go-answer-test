package main

import (
	"fmt"
)

type Book struct {
	Title     string
	Author    string
	Year      int
	Available bool
}

func (b Book) Describe() {
	availability := "No"
	if b.Available {
		availability = "Yes"
	}
	fmt.Printf("Title: %s, Author: %s, Year: %d, Available: %s\n", b.Title, b.Author, b.Year, availability)
}

var library []Book

func AddBook(book Book) {
	library = append(library, book)
}

func FindBook(title, author string) {
	for _, book := range library {
		if book.Title == title && book.Author == author {
			book.Describe()
			return
		}
	}
	fmt.Println("Book not found.")
}

func BorrowBook(title, author string) {
	for i, book := range library {
		if book.Title == title && book.Author == author && book.Available {
			library[i].Available = false
			fmt.Println("Book borrowed successfully.")
			return
		}
	}
	fmt.Println("Book cannot be borrowed.")
}

func DisplayLibrary() {
	for _, book := range library {
		book.Describe()
	}
}

func main() {
	// Add some books to the library
	AddBook(Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: 2015, Available: true})
	AddBook(Book{Title: "Learning Go", Author: "Miek Gieben", Year: 2021, Available: true})

	// Display all books
	fmt.Println("Library Collection:")
	DisplayLibrary()

	// Find a book
	FindBook("Learning Go", "Miek Gieben")

	// Borrow a book
	BorrowBook("The Go Programming Language", "Alan A. A. Donovan")
	
	// Try to borrow the same book again
	BorrowBook("The Go Programming Language", "Alan A. A. Donovan")

	// Display all books again to see the updated status
	fmt.Println("\nUpdated Library Collection:")
	DisplayLibrary()
}
