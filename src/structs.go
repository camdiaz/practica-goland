package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra forma de declarar un struct
	var anotherCar car
	anotherCar.brand = "Ferrari"
	fmt.Println(anotherCar)

}
