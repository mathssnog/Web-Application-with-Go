package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "jogn@gorila.com", 30})
	users = append(users, User{"Mary", "Lory", "maryn@gorila.com", 20})
	users = append(users, User{"Carry", "Anderson", "carry@gorila.com", 50})
	users = append(users, User{"Mat", "Nuiy", "mat@gorila.com", 24})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
