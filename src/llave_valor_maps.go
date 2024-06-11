package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Luisma"] = 27
	m["Cam"] = 23

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["Luisma"]
	fmt.Println(value, ok)
}

// son mas efectivos ya que aplican concurrencia a comparacion de los arrays y los slices en cada una de sus operaciones
