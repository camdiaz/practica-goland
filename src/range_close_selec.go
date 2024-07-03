package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	// len es para ver cuantas goroutines en ese channel
	// cap es para saber cuantas soporta el channel

	fmt.Println(len(c), cap(c))

	// range es el recorrido de cada uno de los valores del canal
	// range y close
	// close cierra los canales apenas se dejan de usar por lo que no recibe mas datos apesar de que tenga capacidad para hacerlo
	close(c)
	// c <- "Mensaje3"

	for message := range c {
		fmt.Println(message)
	}
	// buena forma de iterar entre valores para revisar que es lo que hay dentro del canal

	// select es para escuchar multiples canales pero sabiendo cual funciona primero
	email1 := make(chan string)
	email2 := make(chan string)
	go message("Mensaje1", email1)
	go message("Mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Recibido en email1", m1)
		case m2 := <-email2:
			fmt.Println("Recibido en email2", m2)
		}
	}
}
