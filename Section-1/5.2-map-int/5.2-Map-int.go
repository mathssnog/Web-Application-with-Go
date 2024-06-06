package main

import "log"

func main() {

	myMap := make(map[string]int)
	// É desta forma que declaramos um 'map' através de um make(map);
	// O primeiro índice dentro do colchetes é o índice que utilizamos para procurar o mapa
	// O segundo é o valor armazenado nesse índice, então nesse caso estamos utilizando string's.

	myMap["First"] = 1
	myMap["Second"] = 2

	log.Println(myMap["First"], myMap["Second"])

}
