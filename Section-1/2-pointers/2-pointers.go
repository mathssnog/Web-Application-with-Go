package main

import "log"

func main() {

	// Definição de uma string e fazendo ela receber um valor
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)
	log.Println("After my func call - myString is set to", myString)

}

func changeUsingPointer(s *string) {
	// Essa seria a posição de memória que esse ponteiro estaria apontando
	log.Println("s is set to", s)

	newValue := "red"

	// Fazemos esse local de memória receber o valor que definimos, no caso a string "red"
	*s = newValue
}

// Ponteiro aponta para um local específico da memória retornando aquele local específico de memória.
