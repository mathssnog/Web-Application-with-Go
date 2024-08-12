package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// Corrigimos o RenderTemplate adicionando o começo com letra maiúscula para utilizarmos em outros arquivos.
func RenderTemplate(w http.ResponseWriter, html string) {
	// Precisamos renderizar aqueles arquivos html
	// vamos parsear o template:
	/* Criamos uma variável 'parsedTemplate' o qual recebe o package 'template' que faz o 'template.ParseFiles'
	O que o template.ParseFile faz?
	Ele carrega o arquivo ./templates/ e pega o argumento que eu quero passar como string */

	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get request template from cache

	t, ok := tc[html]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template: ", err)
	// }
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// O text/template é uma lib para manipular templates de texto.

// Aula 31- Building a simple template cache

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var html *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		log.Println("creating template and adding to cache")
// 		// need to create the template
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// have the template in the cache
// 		log.Println("using cached template")
// 	}

// 	html = tc[t]
// 	err = html.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}
// 	// parse the template:
// 	html, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache (map)
// 	tc[t] = html
// 	return nil
// }
