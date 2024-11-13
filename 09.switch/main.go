package main

import "fmt"

func main() {

	day := 3
	//Basic switch
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")
	}

	// Switch with multiple values

	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other session")
	}

	//swith with expressions

	temperature := 10

	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}
