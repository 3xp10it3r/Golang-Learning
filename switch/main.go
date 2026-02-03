package main

import "fmt"


func main() {
	day := 5

	switch day {
		case 1,7,9,5: 
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		default:
			fmt.Println("Default")
	}

	temperature := 20

	switch {
		case temperature < 0:
			fmt.Println("cold")
		case temperature >= 10 && temperature <= 20:
			fmt.Println("Warm")
		default:
			fmt.Println("default")
	}
}