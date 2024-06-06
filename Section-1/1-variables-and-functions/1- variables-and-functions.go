package main

import "fmt"

func main() {

	fmt.Println("Hello World!")

	//parte-1:
	// Vamos agora definir as variáveis:
	// Definimos uma variável do tipo string, fazemos ela receber uma string e printamos essa variável, logo vamos printar a string.
	var whatToSay string
	whatToSay = "good by cruel world"
	fmt.Println(whatToSay)

	//parte-2:
	//Vamos agora definir outro tipo de variável

	var i int
	i = 7
	println("i is", i, "set to")

	// Podemos chamar a função que definimos fora da main:

	whatWassaid, otherThing := saySomething()
	fmt.Println("the function said", whatWassaid, otherThing)
}

//parte-3:
// Vamos agora definir uma função fora da main que retorna 2 strings:
// Lembrando que uma boa prática é termos o nome da função em camelCase

func saySomething() (string, string) {
	return "Something", "else"
}

// Comando para rodar o código: go run course-example-1.go
