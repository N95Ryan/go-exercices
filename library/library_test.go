// library_test.go
package main

import "testing"

func TestAddBook(t *testing.T) {
	lib := Library{}
	book := Book{Title: "Test", Author: "Moi", ID: 1, IsAvailable: true}

	lib.AddBook(book)

	if len(lib.Books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(lib.Books))
	}
	if lib.Books[0].Title != "Test" {
		t.Errorf("Expected book title 'Test', got '%s'", lib.Books[0].Title)
	}
}
