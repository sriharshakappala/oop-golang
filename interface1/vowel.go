package main

import "fmt"

// VowelsFinder is an interface with FindVowels function
type VowelsFinder interface {
	FindVowels() []rune
	// FindFirstVowel() string
}

// CustomString is extended version of string
type CustomString string

// FindVowels finds the vowels and returns the array of vowels in it
func (cs CustomString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range cs {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := CustomString("Sri Harsha Kappala")
	fmt.Printf("Vowels are %c", name.FindVowels())
	var v VowelsFinder
	v = name
	fmt.Println()
	fmt.Printf("Vowels are %c", v.FindVowels())
}
