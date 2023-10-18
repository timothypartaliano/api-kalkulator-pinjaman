package answer

import (
	"fmt"
	"strings"
)

func reverseSentenceAndWords(sentence string) string {
	words := strings.Fields(sentence)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-i-1] = reverseString(word)
	}

	reversedSentence := strings.Join(reversedWords, " ")
	return reversedSentence
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Q6() {
	input := "this is a test"
	reversed := reverseSentenceAndWords(input)
	fmt.Println("Input:", input)
	fmt.Println("Output:", reversed)
}