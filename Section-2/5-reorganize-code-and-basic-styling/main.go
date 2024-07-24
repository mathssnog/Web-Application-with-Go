// let's see what we can serve a real web page
// we gonna reuse the code that we developed on section 3
package main

import (
	"fmt"
	"net/http"
)

// 'const' são utilizadas para variáveis que não vão ter seus valores alterados. Não pode ser mudada por nenhuma parte da aplicação.
const portNumber = ":8080"

func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

// precisamos realizar o comando para rodar: go run *.go
