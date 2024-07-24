// let's see what we can serve a real web page
// we gonna reuse the code that we developed on section 3
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// 'const' são utilizadas para variáveis que não vão ter seus valores alterados. Não pode ser mudada por nenhuma parte da aplicação.
const portNumber = ":8080"

// Lembrando que quando criamos uma função que lida com os parâmetros do browser precisamos de 2 parâmetros: request and response
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, html string) {
	// Precisamos renderizar aqueles arquivos html
	// vamos parsear o template:
	/* Criamos uma variável 'parsedTemplate' o qual recebe o package 'template' que faz o 'template.ParseFiles'
	O que o template.ParseFile faz?
	Ele carrega o arquivo ./templates/ e pega o argumento que eu quero passar como string */
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
