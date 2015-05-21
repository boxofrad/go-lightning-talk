package main

import "fmt"

func main() {
	name, age := getDetails()
	fmt.Println("Hi", name, "you are", age, "years old")
}

func getDetails() (string, int) {
	return "Daniel", 22
}
