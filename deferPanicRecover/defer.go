package deferPanicRecover

import (
	"fmt"
	"log"
)

// defer es una palabra clave que se usa para posponer la ejecución de una función hasta que la función circundante regrese.
func ShowDefer() {
	fmt.Println("Hello")
	defer fmt.Println("Bye jejejjeje")
	defer fmt.Println("World")
	fmt.Println("Bye")

}

// panic es una función que detiene la ejecución de la función actual.
// recover es una función que se usa para recuperar el control de una goroutine en pánico.
func ShowPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Has ocurred an d error by panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("hubo un error")
	}
}
