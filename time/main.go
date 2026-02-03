package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of current time is %T\n", currentTime)

	formatted := currentTime.Format("2006-01-02, 15:04:05")
	fmt.Println(formatted)

	formatted2 := currentTime.Format("2006/01/02, Monday, 3:04 PM")
	fmt.Println(formatted2)

	layoutStr := "02/01/2006"
	dateStr := "31/01/2026"
	
	formatted_time, _ := time.Parse(layoutStr, dateStr)
	fmt.Println(formatted_time)

	// add 1 more day in current time
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new_date time: ", new_date)
 
	formatted_new_date := new_date.Format("2006/01/02, Monday")
	fmt.Println("new date time: ", formatted_new_date)
}