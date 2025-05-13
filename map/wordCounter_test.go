package main

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Texte simple",
			input: "bonjour bonjour ryan",
			expected: map[string]int{
				"bonjour": 2,
				"ryan":    1,
			},
		},
		{
			name:  "Texte avec ponctuation",
			input: "Bonjour, bonjour! Ryan... Ryan!",
			expected: map[string]int{
				"bonjour": 2,
				"ryan":    2,
			},
		},
		{
			name:  "Texte avec espaces multiples",
			input: "  bonjour   bonjour  ryan  ",
			expected: map[string]int{
				"bonjour": 2,
				"ryan":    1,
			},
		},
		{
			name:     "Texte vide",
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.input)

			// Vérifier que tous les mots attendus sont présents avec le bon compte
			for word, expectedCount := range tt.expected {
				if count, exists := result[word]; !exists {
					t.Errorf("Mot manquant: %s", word)
				} else if count != expectedCount {
					t.Errorf("Pour le mot '%s', attendu %d occurrences, obtenu %d",
						word, expectedCount, count)
				}
			}

			// Vérifier qu'il n'y a pas de mots supplémentaires
			for word, count := range result {
				if _, exists := tt.expected[word]; !exists {
					t.Errorf("Mot inattendu trouvé: %s avec %d occurrences", word, count)
				}
			}
		})
	}
}
