package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error getting Get response", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Type of response: %T\n", resp)
	// fmt.Println(resp)

	// Read the response body

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response :", err)
		return
	}

	fmt.Println("Response : ", string(data))
}
