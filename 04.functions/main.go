package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")

}

func add(a, b int) int { // common int for multiple varibles
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learning function in golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("add of two number is:", ans)

	data := multiply(3, 4)
	fmt.Println("multiplication of two number is :", data)

}
