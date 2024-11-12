package main

import "fmt"

func main() {

	x := 5

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 {

		fmt.Println("x is lesser than 5")
	}

	z := 10

	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than 10")
	} else {
		fmt.Println("z is smaller than 5")
	}

	y := 10

	if x < 5 && (y > 5 || z < 40) {
		fmt.Println("how are you doing today ?")
	} else {
		fmt.Println("Learn programming with Hello World")
	}

}
