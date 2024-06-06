package main

import "fmt"

// ////////////////////////////////////////////////////////////////////////////////////////////
// USO DE FUNCIONES
func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// Que me retorne un valor
func returnValue(a int) int {
	return a * 2
}

// Que me retorne doble valor o dos valores al tiempo
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello World!")
	tripleArgument(1, 2, "Hola")

	// retorno del valor
	value := returnValue(2)
	fmt.Println("Value:", value)

	// retorno de dos valores
	value1, value2 := doubleReturn(2)
	fmt.Println("Value1 y Value2:", value1, value2)

	// Si queremos omitir un valor cuando retorna dos o mas valores
	value1, _ = doubleReturn(2)
	fmt.Println("Value1:", value1)
}
