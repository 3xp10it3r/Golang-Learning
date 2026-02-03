package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int `json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main() {
	person := Person{
		Name: "Digvijay",
		Age: 25,
		IsAdult: true,
	}

	fmt.Println(person)

	// marshalling -> encoding to json
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error while Marshalling", err)
		return
	}

	fmt.Println("JSON Data : ", string(jsonData))

	// UnMarshalling -> Decoding a json
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("error unmarshalling", err)
		return
	}

	fmt.Println("Unmarshalled Json :", newPerson)
}
