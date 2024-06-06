package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	// Podemos colocar assim: 'myMap := make(map[string]interface())' --- podemos armazenar o que quisermos. Porém não é recomendado.

	me := User{
		FirstName: "Matheus",
		LastName:  "Nogueira",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	// ---
	// Os maps não estão ordenados pela ordem que você adiciona, elas são programas no sistemas e não na ordem que você adiciona
}
