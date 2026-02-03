package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

func add(a, b int) (result int) {
	result = a + b;
	return
}

func main() {
	simpleFunction()

	ans := add(3, 4)

	fmt.Println("Add two ", ans)
}
