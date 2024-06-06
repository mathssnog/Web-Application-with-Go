package main

import "log"

func main() {

	myVar := "pupu"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	case "fish":
		log.Println("cat is set to fish")

	default: //se nenhuma das condições for true, faça isso:
		log.Println("cat is somthing else")
	}

}
