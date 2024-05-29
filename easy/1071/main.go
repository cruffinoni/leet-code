package main

import (
	"log"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	// Trouver la longueur minimum des deux chaînes
	minLength := min(len(str1), len(str2))

	// Vérifier les sous-chaînes de longueur décroissante
	for length := minLength; length > 0; length-- {
		if len(str1)%length == 0 && len(str2)%length == 0 {
			candidate := str1[:length]
			if isDivisible(str1, candidate) && isDivisible(str2, candidate) {
				return candidate
			}
		}
	}
	return ""
}

// Fonction pour vérifier si une chaîne est formée par des répétitions d'une sous-chaîne
func isDivisible(str, sub string) bool {
	repeated := strings.Repeat(sub, len(str)/len(sub))
	return repeated == str
}

type test struct {
	str1 string
	str2 string
}

func main() {
	tests := map[string]test{
		"ABC":   {str1: "ABCABC", str2: "ABC"},
		"AB":    {str1: "ABABAB", str2: "ABAB"},
		"":      {str1: "ABCDEF", str2: "ABC"},
		"TAUXX": {str1: "TAUXXTAUXXTAUXXTAUXXTAUXX", str2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"},
	}
	for expected, t := range tests {
		res := gcdOfStrings(t.str1, t.str2)
		if res != expected {
			log.Printf("expected '%s', got '%s'", expected, res)
		} else {
			log.Printf("test '%s' passed!", expected)
		}
	}
}
