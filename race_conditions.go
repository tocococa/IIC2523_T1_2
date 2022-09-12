package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var number float64 = 1

func main() {
	/* Función que llama primero a inc y luego a square, donde primero se suma 1
	a number y luego se eleva al cuadrado, donde el resultado esperado es 4.
	Sin embargo, y debido a la falta de sincronización para acceder a la variable
	si el calculo de inc es mas lento que el de square, primero se eleva 1 al cuadrado y
	luego se suma, dando como resultado 2.
	*/
	wg.Add(2)
	go inc()
	go square()
	wg.Wait()
	fmt.Println(number)
}

func inc() {
	/* Primero hace algunos calculos pesados que dependen de una fuente
	no determinista, y luego suma 1 a number
	*/
	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := float64(rand.New(source).Intn(100))
	// Se hacen unos calculos pesados
	counter := 2
	for i := 0; i < int(math.Pow(randomNumber, 2)); i++ {
		counter = counter * 2
	}
	// se le suma 1 a number
	number += 1
	wg.Done()
}

func square() {
	/* Primero hace algunos calculos pesados que dependen de una fuente
	no determinista, y luego eleva number al cuadrado
	*/
	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := float64(rand.New(source).Intn(100))
	// Se hacen unos calculos pesados
	counter := 2
	for i := 0; i < int(math.Pow(randomNumber, 2)); i++ {
		counter = counter + 1
	}
	// se eleva number al cuadrado
	number = math.Pow(number, 2)
	wg.Done()
}
