package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, World!")
}

func sayHello() {
	fmt.Println("Say Hello")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Niiii")
}

func main() {
	fmt.Println("Learning Goroutines..")
	go sayHello()
	hello()

	time.Sleep(1000 * time.Millisecond)
}
