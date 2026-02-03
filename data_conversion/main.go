package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42

	fmt.Println("Number is: ", num)
	fmt.Printf("Type of num is %T\n", num)

	data := float64(num)
	data += 1.23
	fmt.Println("Number is: ", data)
	fmt.Printf("Type of num is %T\n\n", data)

	nnum := 123
	str := strconv.Itoa(nnum)
	fmt.Println("str is: ", str)
	fmt.Printf("Type of str is %T\n\n", str)


	number_string := "58493"
	number_int, _ := strconv.Atoi(number_string)
	number_int += 2
	fmt.Println("number_int is: ", number_int)
	fmt.Printf("Type of number_int is %T\n\n", number_int)

	PI_string := "3.14"
	PI_float, _ := strconv.ParseFloat(PI_string, 64)
	fmt.Println("PI is", PI_float)
	fmt.Printf("Type of PI is %T\n\n", PI_float)


}
