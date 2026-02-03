package main

import "fmt"

func main() {
	// name -> grade

	studentsGrades := make(map[string]int)

	studentsGrades["Happy"] = 100
	studentsGrades["Digvijay"] = 100
	studentsGrades["Any"] = 10

	fmt.Println("Marks of Happy : ", studentsGrades["Happy"])
	fmt.Println("Marks of Any : ", studentsGrades["Any"])


	delete(studentsGrades, "Any")
	fmt.Println("Marks of Any : ", studentsGrades["Any"])

	grades, exists := studentsGrades["Any"]

	fmt.Printf("Grades of any : %d\nAny Exists ? -> %v\n", grades, exists)

	for index, value := range studentsGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	persons := map[string]int{
		"Alice": 90,
		"Bob": 85,
		"Charlie": 10,
	}

	for index, value := range persons {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}
}
