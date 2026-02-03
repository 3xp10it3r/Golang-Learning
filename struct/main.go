package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	City string
	Zip int
}

type Employee struct {
	person Person
	contact Contact
	address Address
}

func main() {
	var happy Person
	fmt.Println("Happy person: ", happy)

	happy.FirstName = "Happy"
	happy.LastName = "Gupta"
	happy.Age = 24

	fmt.Println("Happy person: ", happy)

	//2nd method
	person1 := Person {
		FirstName: "Digvijay",
		LastName: "Gupta",
		Age: 25,
	}

	fmt.Println("Person 1: ", person1)


	var person2 = new(Person)
	person2.FirstName = "Simran"
	person2.LastName = "Agarwal"
	person2.Age = 26

	fmt.Println(person2)
	fmt.Println(person2.FirstName, person2.LastName)

	fmt.Println("-------------------------")

	var emp Employee

	emp.person = Person{
		FirstName: "Dig",
		LastName: "Vijay",
		Age: 25,
	}

	emp.contact = Contact{
		Email: "dig@gmail.com",
		Phone: "752398373",
	}

	emp.address = Address{
		City: "Kza",
		Zip: 274802,
	}

	fmt.Println(emp)
	fmt.Printf("%s%s\n",emp.person.FirstName, emp.person.LastName)

}
