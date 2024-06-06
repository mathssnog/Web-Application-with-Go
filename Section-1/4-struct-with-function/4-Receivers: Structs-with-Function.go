package main

import "log"

type myStruct struct {
	FirstName string
}

func main() {

	var myVar myStruct

	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}
	//Obs.: Lembrando que '=' é para atribuição de um valor quando a variável já está declarada
	// e ':=' é para atribuição também mas quando não temos a variável declarada.

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar is set to", myVar2.FirstName)
}
