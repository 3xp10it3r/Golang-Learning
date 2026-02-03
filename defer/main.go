package main

import "fmt"

func add(a,b int) int {
	return a + b
}

func main() {
	fmt.Println("start of the program")
	data := add(5,6)
	defer fmt.Println("data is : ", data)
	defer fmt.Println("middle of the program")
	fmt.Println("end of the program")

}

// defer pushes things in stack in order, so when it works in reverse way.
