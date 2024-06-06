package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With And
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es 1 y 2")
	} else {
		fmt.Println("No es 1 y 2")
	}

	// With Or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	} else {
		fmt.Println("No es 0 o 2")
	}

	// Convert text to numeric
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}
