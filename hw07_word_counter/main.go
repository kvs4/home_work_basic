package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	sample := "five, five, five one.. two: three/// one"
	wordCounts := countWords(sample)
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWords(text string) map[string]int {
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(text, f)
	wordCounts := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word) // conversion the word to lowercase
		wordCounts[word]++
	}
	return wordCounts
}
