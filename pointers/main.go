package main

import "fmt"

func modifyByReference(num *int) {
	*num = *num + 20
}

func main() {
	var num int = 2

	// var ptr *int
	
	ptr := &num

	fmt.Println("Num has value:", num)
	fmt.Println("pointer contains:", ptr)
	fmt.Println("pointer value:", *ptr)


	var pointer *int // default value -> nil

	if pointer == nil {
		fmt.Println("Pointer is not assigned.")
	}

	value := 10

	modifyByReference(&value)

	fmt.Println(value)

}
