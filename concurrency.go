package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go sayHello(1, done)
	go sayHello(2, done)

	<-done
}

func sayHello(num int, done chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello from goroutine:", num, "- iteration:", i)
	}

	done <- true
}
