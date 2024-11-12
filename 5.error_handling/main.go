package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil

}

func main() {
	fmt.Println("standard Error Handling in the functions")

	ans, _ := divide(10, 0)
	fmt.Println("Division of two numbers is ", ans)

}
