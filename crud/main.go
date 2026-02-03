package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error Getting :", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response :", resp.Status)
		return
	}

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading :", err)
	// 	return
	// }

	// fmt.Println(string(data))

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)

	if err != nil {
		fmt.Println("Error while decoding", err)
		return
	}

	fmt.Println("Todo :", todo)
}

func performPostRequest() {
	todo := Todo {
		UserId: 20,
		Title: "Digvijay gupta",
		Completed: false,
	}

	// convert todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))

	fmt.Println("Res status :", res.Status)

}

func main() {
	performPostRequest()
}
