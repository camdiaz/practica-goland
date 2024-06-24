// Personalizar el output en consola de un struct
package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	// Crea un string con el resultado de la funcion
	// d es int, s es string
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)
}
