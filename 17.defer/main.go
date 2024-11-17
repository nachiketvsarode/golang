// in all programing languages, all lines are executed top to bottom linearnely
// if we use defer before any statement or fucntion then it will run at last
// if there are multiple defer statements then beased on order of execution - LIFO / Stack rules will be followed.

package main

import "fmt"

func main() {

	fmt.Println("First line of statement")
	fmt.Println("Second line of statement")
	fmt.Println("Third line of statement")

}
