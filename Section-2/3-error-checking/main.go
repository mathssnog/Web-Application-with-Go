package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 'const' são utilizadas para variáveis que não vão ter seus valores alterados. Não pode ser mudada por nenhuma parte da aplicação.
const portNumber = ":8080"

// Lembrando que quando criamos uma função que lida com os parâmetros do browser precisamos de 2 parâmetros: request and response
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!!")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

// Lembrando que quando temos o caso de AddValues começando com maiúsculo, vamos chamar essa função em outros arquivos, se for com minúsculo vamos utilizar somente aqui.
func addValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func Devide(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Devide Page!!")
	f, err := devideValues(100.0, 10.2)
	if err != nil {
		fmt.Fprintf(w, "cannot devide by 0")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 10.2, f))
}

func devideValues(x, y float32) (float32, error) {
	if y <= 0.0 {
		err := errors.New("Cannot devide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/devide", Devide)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
