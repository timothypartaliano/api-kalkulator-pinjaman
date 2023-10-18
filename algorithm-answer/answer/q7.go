package answer

import (
	"fmt"
	"strings"
)

func reverseWordsInSentence(sentence string) string {
	words := strings.Fields(sentence)

	for i, word := range words {
		words[i] = reverseStrings(word)
	}

	reversedSentence := strings.Join(words, " ")
	return reversedSentence
}

func reverseStrings(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Q7() {
	input := "Thankyou for coming"
	reversed := reverseWordsInSentence(input)
	fmt.Println("Input:", input)
	fmt.Println("Output:", reversed)
}
