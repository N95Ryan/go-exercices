package main

import "fmt"

// Création de la struct des livres
type Book struct {
	Title       string
	Author      string
	ID          int
	IsAvailable bool
}

func main() {
	// Création de plusieurs livres
	book1 := Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", ID: 1, IsAvailable: true}
	book2 := Book{Title: "1984", Author: "George Orwell", ID: 2, IsAvailable: false}
	book3 := Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", ID: 3, IsAvailable: true}

	// Création du slice de livres
	library := []Book{book1, book2, book3}

	// Affichage des livres
	for _, book := range library {
		fmt.Printf("Livre: %d\n", book.ID)
		fmt.Printf("Titre: %s\n", book.Title)
		fmt.Printf("Auteur: %s\n", book.Author)
		fmt.Printf("Disponible: %t\n", book.IsAvailable)
	}

}
