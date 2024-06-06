package main

import "log"

func main() {
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 11
	mySecondString := "another string"

	log.Println(myString, mySecondString, myInt)
}

//Poderiamos ter esse tipo de declaração, porém pdemos utilizar o 'Map'
