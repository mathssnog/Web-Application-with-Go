package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{
		FirstName: "trevor",
		LastName:  "sawler",
	}

	log.Println(user.FirstName, user.LastName, "Bithdate:", user.BirthDate)
}

// Comentou sobre as letras maiúsculas quando definimos uma variável:
// Quando temos por exemplo 'var Special string' podemos utilizar em outros repos
// se for 'var special string' não podemos utilizar, ou seja, serve somente para esse repo
