package main

import (
	"fmt"
	"os"
)

func main() {
	/* 
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Successfully created the file")

	content := "hello world by Digvijay Gupta"
	_, errors := io.WriteString(file, content + "\n")
	if errors != nil {
		fmt.Println("Unable to write to the file : ", errors)
		return 
	}

	fmt.Println("Successfully wrote to the file")
	*/


	/*
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()

	// create a buffer
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file", err)
			return
		}

		fmt.Println(string(buffer[:n]))
	}
	*/

	// read entire file into a byte slice
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("error while reading file :" , err)
		return
	}

	fmt.Println(string(content))

}
