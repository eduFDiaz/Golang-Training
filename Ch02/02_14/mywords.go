// Calculate maximal value in a slice
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
		Needles and pins
		Needles and pins
		Sew me a sail
		To catch me the wind
	`
	word := strings.ToLower("To")
	words, wordcount := wordCount(strings.Fields(strings.ToLower(text)), word)
	fmt.Println(text)
	fmt.Println(words)
	fmt.Printf("The word %q was found in the poem %d times", word, wordcount)
}

// This function returns the amount of times a word is found in an array of strings
func wordCount(array []string, word string) (map[string]int, int) {
	counts := map[string]int{}
	for _, word := range array {
		counts[word]++
	}
	return counts, counts[word]
}
