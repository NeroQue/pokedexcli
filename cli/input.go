package cli

import (
	"bufio"
	"io"
	"strings"
)

func CleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	words := strings.Fields(cleaned)

	seen := make(map[string]bool)
	result := make([]string, 0, len(words))

	for _, word := range words {
		cleanedWord := strings.Trim(word, ".,!?")
		if !seen[cleanedWord] {
			seen[cleanedWord] = true
			result = append(result, cleanedWord)
		}
	}
	return result
}

func NewScanner(r io.Reader) *bufio.Scanner {
	return bufio.NewScanner(r)
}
