package main

import "log"

type myStruct struct {
	FirstName string
}

// Colocamos logo após o 'func' o nome de uma variável apontando para a struct 'myStruct'
// O que isso faz ? Isso é chamado de Receiver, e vincula minha struct à função 'printFirstName'
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {

	var myVar myStruct

	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}
	//Obs.: Lembrando que '=' é para atribuição de um valor quando a variável já está declarada
	// e ':=' é para atribuição também mas quando não temos a variável declarada.

	//Chamamos a função 'printFirstName' que é parte da myStruct vinculada através do Receiver.
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar is set to", myVar2.printFirstName())
}
