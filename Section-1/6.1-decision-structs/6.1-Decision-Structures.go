package main

import "log"

func main() {
	//Vamos aprender o statement "if"

	var myNum int
	var isTrue bool

	myNum = 100
	isTrue = false

	if myNum > 99 && !isTrue {
		log.Println("My number is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else if myNum == 101 || isTrue {
		log.Println("2")
	} else if myNum > 1000 && isTrue == false {
		log.Println("3")
	}
}
