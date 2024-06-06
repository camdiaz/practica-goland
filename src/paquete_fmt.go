package main

import "fmt"

func main() {
	// /////////////////////////////////////////////////////////////////////////////////////
	// PAQUETE FMT: ALGO MAS QUE IMPRIMIR EN CONSOLA
	// Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println se usa como salto de linea de forma automatica
	fmt.Println(helloMessage, worldMessage)

	// Printf se usa para imprimir con formato y agrega una funcion extra al string que insertas como valor de entrada
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf genera un string pero no lo impreme en consola
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// TIPOS DE DATOS para saber que tipo de dato es
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
