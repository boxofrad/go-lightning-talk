package main

import "fmt"

type Dog struct {
	Name string
}

func (dog *Dog) Bark() {
	fmt.Printf("Woof! by name is %s", dog.Name)
}

func main() {
	rover := Dog{Name: "Rover"}
	rover.Bark()
}
