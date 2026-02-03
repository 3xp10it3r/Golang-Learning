package main

import "fmt"

func main() {
	var name[5] string

	name[0] = "Digvijay"
	name[1] = "Happy"

	fmt.Println(name)
	fmt.Println("-------")
	fmt.Printf(" %q\n",name)

	var numbers = [8]int{1,2,3,4,5}

	fmt.Println(numbers, len(numbers))
}