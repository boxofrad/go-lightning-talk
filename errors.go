package main

import "os"

func main() {
	file, err := os.Open("file.go")

	if err != nil {
		// handle the error
	}
}
