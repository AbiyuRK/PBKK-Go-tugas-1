package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hasMinWords(input string, minWords int) bool {
	words := strings.Fields(input)
	return len(words) >= minWords
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func main() {
	input := getUserInput()
	const minWords = 3

	if hasMinWords(input, minWords) {
		fmt.Println("Reversed string:", reverseString(input))
	} else {
		fmt.Printf("Input must contain at least %d words.\n", minWords)
	}
}
