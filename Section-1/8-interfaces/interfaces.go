package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
	NumberOfTeeth() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name  string
	Color string
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	printInfo(&dog)

	gorilla := Gorilla{
		Name:  "Jock",
		Color: "grey",
	}

	printInfo(&gorilla)
}

func printInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs and have", a.NumberOfTeeth(), "tooths")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Dog) NumberOfTeeth() int {
	return 36
}

func (d *Gorilla) Says() string {
	return "woof"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func (d *Gorilla) NumberOfTeeth() int {
	return 38
}
