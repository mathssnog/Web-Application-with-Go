package main

import "fmt"

func main() {
	fmt.Println("Starting function")
	defer fmt.Println("Ending function") // Adia a execução até o fim da função

	fmt.Println("Function in progress")
	// Mais código
	// ...
}

// O output desse cara é primeiro:
/*
	- "starting function"
	- "function in progress"
	- "ending function"

Ou seja, a função é inteiramente executada, e após ela terminar, temos a execução do 'defer'. Então ele praticamente adia a execução do argumento ao qual
foi designado até o término da função.
*/
