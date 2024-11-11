package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Learn go language by hello world")

	myutil.PrintMessage("Hello World")

	var name string = "Prince"
	var version = 26
	var currency = 1234567
	fmt.Println(name)
	fmt.Println(version)
	fmt.Println("currency: ", currency)

	var dimension float64 = 87.12
	fmt.Println(dimension)

	var decision bool = true
	fmt.Println(decision)

	const pi = 3.14
	fmt.Println(pi)

	//shorthand

	person := "Nachiket"
	fmt.Println(person)

	var Publicvariable = "data is public" // check the uppercase "p" in variable public, if variable declared with ffirst character uppercase then it can be exported

	var privatevariable = "data is private" // lowercase first char of variable cannot be exported

	fmt.Println(Publicvariable)
	fmt.Println(privatevariable)

}

func Publicfunction() {
	// same rule of P and p

}

func privatefunction() {
	// same rule of P and p

}
