package main

import (
	"fmt"
	"strings"
	"unicode"
)

// CountWords compte les occurrences de chaque mot dans un texte
func CountWords(text string) map[string]int {
	// Création d'un map pour compter les mots
	wordMap := make(map[string]int)

	// Fonction pour nettoyer un mot
	cleanWord := func(word string) string {
		// Convertir en minuscules
		word = strings.ToLower(word)

		// Supprimer la ponctuation
		return strings.Map(func(r rune) rune {
			if unicode.IsPunct(r) {
				return -1 // Supprime le caractère
			}
			return r
		}, word)
	}

	// Diviser le texte en mots en gérant les espaces multiples
	words := strings.Fields(text)

	// Compter les mots
	for _, word := range words {
		// Nettoyer le mot
		clean := cleanWord(word)

		// Ignorer les mots vides
		if clean != "" {
			wordMap[clean]++
		}
	}

	return wordMap
}

func main() {
	// Exemple d'utilisation
	text := "Bonjour bonjour Ryan, Ryan ! Go Go! go..."
	result := CountWords(text)

	// Affichage des résultats
	fmt.Println("Nombre d'occurrences de chaque mot :")
	for word, count := range result {
		fmt.Printf("'%s' : %d fois\n", word, count)
	}
}
