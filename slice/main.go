package main

import "fmt"

func main() {
	numbers := []int{1,2,3}

	numbers = append(numbers, 4,5,6)

	var names = []string{}

	fmt.Printf("type : %T\n", numbers)

	fmt.Println(names, numbers, len(numbers), cap(numbers))

	// using make([]type, length, capacity)

	nums := make([]int, 2, 5)

	nums = append(nums, 1,3,4)

	fmt.Println("slice: ", nums)
	fmt.Println("len: ", len(nums))
	fmt.Println("capacity: ", cap(nums))

	nums = append(nums, 6)

	fmt.Println("slice: ", nums)
	fmt.Println("len: ", len(nums))
	fmt.Println("capacity: ", cap(nums))

// slice:  [0 0 1 3 4 6]
// len:  6
// capacity:  10

}