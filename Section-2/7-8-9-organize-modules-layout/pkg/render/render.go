package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Corrigimos o RenderTemplate adicionando o começo com letra maiúscula para utilizarmos em outros arquivos.
func RenderTemplateTest(w http.ResponseWriter, html string) {
	// Precisamos renderizar aqueles arquivos html
	// vamos parsear o template:
	/* Criamos uma variável 'parsedTemplate' o qual recebe o package 'template' que faz o 'template.ParseFiles'
	O que o template.ParseFile faz?
	Ele carrega o arquivo ./templates/ e pega o argumento que eu quero passar como string */
	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var html *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache")
		// need to create the template
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// have the template in the cache
		log.Println("using cached template")
	}

	html = tc[t]
	err = html.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}
	// parse the template:
	html, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add template to cache (map)
	tc[t] = html
	return nil
}
