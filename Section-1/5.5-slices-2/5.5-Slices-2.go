package main

import "log"

func main() {

	numbers := []int{1, 3, 4, 5, 76, 85, 32}

	log.Println(numbers)
	log.Println(numbers[0:2]) // somente os 2 primeiros numeros
}

// podemos fazer o mesmo sistema para strings.
