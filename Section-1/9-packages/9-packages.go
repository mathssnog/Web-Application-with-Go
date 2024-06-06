package main

import "log"

func main() {
	log.Println("HELLO!")

}

// Comentou sobre go modules, e que antigamente precisavamos definir uma variável de ambiente $GOPATH e colocar as
// coisas em determinadas pastas.

// Abriu o terminal: go mod init <github-name>
// No caso ele iniciou o módulo: go mod init github.com/tsawler/myniceprogram

// Convencionalmente, você começa a nomear as coisas prefixando-as com o nome do repositório git onde elas ficarão.
// Podemos chamá-lo como quiser, mas é útil usar essa convenção de nomenclatura, mesmo que não faça push para o git.

// Posteriormente a isso ele fez o seguinte:
// Golang preferencies e depois procurou por Go Modules e fez o seguinte: Ativou a flag - Enable Go Modules Integration
/* 		Fez isso:

package
	|__helpers
	|	|
	|	|__helpers.go
	|
	|__go.mod
	|__main.go

Dentro do helpers.go temos:
---
package helpers

type SomeType struct{
	TypeName string
	TypeNumber int
}
---

Ou seja, criamos nosso próprio package.

Comentou que podemos utilizar agora essa struct no main.go:
---
package main

import {
	"log"
	"github.com/tsawler/myniceprogram/helpers"
}


func main() {
	log.Println("HELLO!")

	var myVar helpers.SomeType
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}
---

Ou seja conseguimos utilizar a struct do helpers na main, utilizando essa parada de packages. */
