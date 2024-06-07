package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //a função de tratamento exige que a solicitação seja um ponteiro.
		n, err := fmt.Fprintf(w, "Hello world") // Esse cara aqui está mostrando no browser quando acessamos nosso localhost.
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)

}
