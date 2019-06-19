package interfaceTest

import (
	"fmt"
	"testing"
)

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune

	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func TestFindVowels(t *testing.T) {
	// name := MyString("Sam Anderson")

	var name MyString = "Sam Anderson"

	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
}
