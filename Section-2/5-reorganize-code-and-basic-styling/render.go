package main

import (
	"fmt"
	"net/http"
	"text/template"
)

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
