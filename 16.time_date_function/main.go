package main

import (
	"fmt"
	"time"
)

//golang has fixed time and date format in 24 hour format
// eg. 2006-01-02 15:04:05
//YYYY : 2006
//MM : Month
//DD : date
//hh: hours in 24 hours format
//mm: minutes
//ss: seconds

func main() {

	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of CurrentTime %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006, 15:04:05")
	fmt.Println("Formatted time: ", formatted)

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time:", formatted_time)

	// add 1 more day in currentTime

	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new_date time:", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_date time:", formatted_new_date)

}

/*
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
*/
