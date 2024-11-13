package main

import "fmt"

func main() {

	fmt.Println("We are learning arry in golang")

	var name [5]string
	name[0] = "Nachiket"
	name[2] = "Raj"

	fmt.Println("Name of Person is :", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is:", numbers)

	fmt.Println("Length of numbers is :", len(numbers))

	fmt.Println("value of name at index 2 is:", name[2])

	var price [5]int
	fmt.Println("Price is :", price) //default value of int, float is 0; string is "", boolean is false, pointer and complex is nil

	var city [5]string
	city[0] = "Toronto"
	city[2] = "York"
	fmt.Printf("cost array is %q\n", city) // %q is for the quoted string

}
