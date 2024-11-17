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

	formatted := currentTime.Format("dd-mm-yyyy")
	fmt.Println("Formatted time: ", formatted)

}
