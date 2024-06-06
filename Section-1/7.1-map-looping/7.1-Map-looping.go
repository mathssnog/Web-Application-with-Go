package main

import "log"

func main() {
	animals := make(map[string]string)

	animals["dog"] = "Fido"
	animals["horse"] = "Fluffy"

	for animalType, animals := range animals {
		log.Println(animalType, animals)
	}

	// Aqui podemos ter o índice sendo um parâmetro valioso, porque temos o tipo do animal sendo computado no map
}
