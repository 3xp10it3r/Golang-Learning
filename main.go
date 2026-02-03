package main

import (
	"fmt"
	"helloworld/myutil"
)

func main() {
	fmt.Println("Learn go langauge from Digvijay - Hello World!")
	myutil.PrintMessage("----Hey---")

	var name string = "Digvijay"
	var version = 1

	fmt.Println(name)
	fmt.Println(version)

	var money int = 67000
	fmt.Println(money)

	var dimention float64 = 87.902
	fmt.Println("Dimention: ", dimention)

	var decided bool = true
	decided = false;

	fmt.Println(decided)

	const PI = 3.14
	fmt.Println(PI)

	person := "Happy Gupta"
	fmt.Println(person)

	var PublicVariable = "Data is public so start with capital letter"
	var privateVariable = "Data is public so start with small letter"

	fmt.Println(PublicVariable, " ", privateVariable)
}