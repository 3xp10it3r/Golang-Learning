package main

import "fmt"

func main() {
	age := 25
	name := "Happy"
	height := 5.8493039

	fmt.Println("age: ", age, "Name : ", name, "height : ", height)

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.3f\n", height)
	fmt.Printf("Name is %s\n", name)

	fmt.Printf("Type of height is %T\n", height)
	
}
