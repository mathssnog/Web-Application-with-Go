package main

import (
	"log"
	"sort"
)

func main() {
	// Comentou que em Golang não usamos Arrays, utilizamos Slices

	var mySlice []string //Se colocarmos 2 colchetes transformamos em Slices.
	var mySlice2 []int

	mySlice = append(mySlice, "Matheus") //Lembrando que o append ele coloca no fim do Slice
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	//--- Vamos agora ver o slice de 'int', que é basicamente a mesma coisa:
	mySlice2 = append(mySlice2, 2) //Lembrando que o append ele coloca no fim do Slice
	mySlice2 = append(mySlice2, 10)
	mySlice2 = append(mySlice2, 6)

	log.Println(mySlice)
	log.Println(mySlice2)

	sort.Ints(mySlice2)
	log.Println(mySlice2)
}
