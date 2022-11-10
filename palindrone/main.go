package main

import (
	"fmt"
	"strings"
)

func palindrome(word string) bool {
	reverseString := ""
	for i := len(word) - 1; i >= 0; i-- {
		reverseString += string(word[i])
	}

	return strings.ToLower(string(reverseString)) == strings.ToLower(word)
}

func main() {
	var input string

	fmt.Print("Please input your word: ")
	fmt.Scanf("%s", &input)

	if palindrome(input) {
		fmt.Println("The word is palindrome")
	} else {
		fmt.Println("The word is not palindrome")
	}

}
