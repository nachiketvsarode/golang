// in all programing languages, all lines are executed top to bottom linearnely
// if we use defer before any statement or fucntion then it will run at last
// if there are multiple defer statements then beased on order of execution - LIFO / Stack rules will be followed.

package main

import "fmt"

func add(a, b int) int {
	return a + b

}

func main() {

	fmt.Println("First line of statement")
	defer fmt.Println("Second line of statement")
	data := add(5, 5)
	defer fmt.Println("Data is :", data)
	fmt.Println("Third line of statement")

}
