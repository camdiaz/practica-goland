package main

import (
	"fmt"

	"./mypackage"
)

func main() {
	var myCar mypackage.CarPulic
	myCar.Brand = "Ford"
	fmt.Println(myCar)
}
