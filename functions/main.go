package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")

}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("we are learning function in golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("add of two number is:", ans)

}
