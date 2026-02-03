package main

import "fmt"


func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 0

	for {
		fmt.Println("counter ", counter)

		counter++
		if counter == 3 {
			break
		}
	}

	numbers := []int{21,232,3543,5,3253,423}

	for index, value := range numbers {
		fmt.Printf("Index %d, value %d\n", index, value)
	}

	data := "Hello, World!"

	for _,value := range data {
		fmt.Printf("%c  ", value)
	}
}