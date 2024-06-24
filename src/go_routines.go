package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wq *sync.WaitGroup) {
	defer wq.Done()
	fmt.Println(text)
}
func main() {
	//sync.WaitGroup nos permitira agrupar las GoRutines y les difinira un tiempo determinado dentro de el hilo del nucleo como concurrencia.
	var wq sync.WaitGroup

	fmt.Println("Hello")
	//Se anexa la informacion de las GoRutines que se agruparan junto con la GoRutine Main.
	wq.Add(1)

	//El tiempo de la ejecucion de una GoRutine Adicional se tomara como un tiempo extra, que la funcion main no esperara.
	go say("World", &wq)
	//Justo despues de que se este ejecutando la funcion say, la funcion main estara sobre la linea de codigo 24.
	wq.Wait()
	//con este atributo le diremos a la funcion main que espere, hasta que el atributo .Done(Linea 10) en la funcion ejecutada en una GoRutine adicional,se ejecute.
	//Una vez ejecutada, la funcion main finalizara.
	fmt.Println("Ha finalizado. :)")

	// Tambien Podremos ejecutar dentro de una GoRutine adicional, funciones anonimas.
	go func(text string) {
		fmt.Println(text)
	}(":) ahora sí, adios.")
	//Este time.sleep le esta dando tiempo a que la GoRutine adicional, que contine la funcion anonima, se ejecute.
	time.Sleep(time.Second * 1)

}
