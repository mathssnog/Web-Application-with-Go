//Channels são um meio de enviar informações de uma parte do seu programa para outra parte do programa.

package main

import (
	"chanels/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(intChan chan int) { //recebe um argumento 'intChan' e é um canal do tipo int e não retorna nada
	// Queremos gerar um número aleatório e fazer coisas com ele. Vamos criar esse random number na helpers.go
	RandomNumber := helpers.RandomNumber(numPool)
	intChan <- RandomNumber
}

func main() {
	//Conseguimos criar os channels de forma semelhante à criação de maps "make" key word:

	intChan := make(chan int)
	defer close(intChan)
	// O que o 'defer' diz é que tudo o que vier depois dessa palavra-chave, defer, será executado assim que a função atual for concluída.

	go CalculateValue(intChan) // estamos chamando a função CalculateValue
	num := <-intChan
	log.Println(num)
}
