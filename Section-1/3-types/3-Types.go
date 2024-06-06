package main

import (
	"log"
)

// Como esta variável está declarada fora de todas as funções, significa que esse valor
// serve para todas as funções, ou seja, se declararmos 's' em qualquer função, então
// vamos ter o valor de "seven"
var s = "seven"

func main() {
	var s2 = "six"

	log.Println(s)
	log.Println("s2 is", s2)

	saySomething("xxx")

}

func saySomething(s3 string) (string, string) {
	// Devemos prestar muito atenção na declaração das variáveis, por exemplo:
	// Temos a variável global 's' declarada, então se tivermos a declaração da mesma variável na função e passarmos um valor para ela
	// vamos estar sobreescrevendo a variável global nessa função. Exemplo:
	// func saySomething(s string) (string, string) {
	//		log.Println("s from the saySomething", s)

	// A variável 's' a ser printada será 'xxx' pois sobreescrevemos o valor da global!

	log.Println("s from the saySomething", s)
	return s3, "World"
}
