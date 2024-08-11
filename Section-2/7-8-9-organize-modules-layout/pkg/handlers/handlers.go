// esse cara será utilizado para os handlers
package handlers

import (
	"net/http"

	"github.com/matssnog/web-application-with-go/pkg/render"
)

// Lembrando que quando criamos uma função que lida com os parâmetros do browser precisamos de 2 parâmetros: request and response
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
