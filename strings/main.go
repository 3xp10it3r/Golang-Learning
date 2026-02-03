package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange,apple,banana,pineapple"

	parts := strings.Split(data, ",")
	fmt.Println(parts)

	count := strings.Count(data, "apple")
	fmt.Println("count is : ", count)

	str := "   Hello, Go! "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "Happy"
	str2 := "Gupta"
	result := strings.Join([]string{str1, str2, "In Google"}, " ")
	fmt.Println("Joined : ", result)
}
