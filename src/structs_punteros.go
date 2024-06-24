package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

// Para acceder a los valores del struct mediante el puntero *
func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

// Para crear funciones, y para llevar una funcion de una  libreria a otra
func main() {
	a := 50
	b := &a
	*b = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	// struct
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	// ram actual
	fmt.Println(myPC)
	myPC.duplicateRAM()

	// ram duplicada

	fmt.Println(myPC)
	myPC.duplicateRAM()

	// ram duplicada x2
	fmt.Println(myPC)

	// De esta forma se puenden usar punteros accediendo a la memoria y pudiendo modificar los datos de forma mas optima

}
