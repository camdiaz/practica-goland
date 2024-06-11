package main

import "fmt"

func main() {
	slice :=
		[]string{"Hola", "Mundo", "Go"}

	for i, value := range slice {
		fmt.Println(i, value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindrome("ama")
}

func isPalindrome(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
