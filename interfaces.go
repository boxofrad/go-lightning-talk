package main

import "fmt"

type Animal interface {
	MakeNoise()
}

type Cat struct{}

func (cat Cat) MakeNoise() {
	fmt.Println("Purr")
}

type Dog struct{}

func (dog Dog) MakeNoise() {
	fmt.Println("Bark")
}

func main() {
	animals := []Animal{Cat{}, Dog{}}
	annoyNeighbours(animals)
}

func annoyNeighbours(animals []Animal) {
	for _, animal := range animals {
		animal.MakeNoise()
	}
}

// end OMIT
