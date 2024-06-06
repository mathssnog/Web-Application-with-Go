package main

import "log"

func main() {
	/* vamos ver loopings
	/for i := 0; i <= 5; i++ {
	/	log.Println(i)
	}
	*/

	animals := []string{"dog", "fish", "cat", "horse", "cow"}

	for _, animal := range animals {
		log.Println(animal)
	}

	// o Valor em branco '_' significa que não estamos interessados nos índices do slice(array). Desta forma vamos ter somente os valores dentro
	// do slice.

}
