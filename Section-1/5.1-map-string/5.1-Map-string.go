package main

import "log"

func main() {

	myMap := make(map[string]string)
	// É desta forma que declaramos um 'map' através de um make(map);
	// O primeiro índice dentro do colchetes é o índice que utilizamos para procurar o mapa
	// O segundo é o valor armazenado nesse índice, então nesse caso estamos utilizando string's.

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "Fido"
	// Se colocarmos o mesmo índice, vamos sobrepor o índice anterior

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

}
