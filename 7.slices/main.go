//slice is a flexible and dynamic datastructure which is alternative to arrays which has a dynamic size as compare to static sizes of arrays,
//Length of slices can be changed during the program's execution.
//The make function allows you to pre-allocate memory for a slice with a specified length and capacity

package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 10, 11, 12, 13, 14, 15)
	fmt.Println("Numebrs:", numbers)

	fmt.Printf("numbers has datatype: %T\n", numbers) //%T is to determine the datatype

	fmt.Println("Length:", len(numbers))
	fmt.Println("Slice:", numbers)
	fmt.Println("Capacity:", cap(numbers))

	units := make([]int, 3, 5) //syntax :make([]T, length, capacity) where T is type of slice
	units = append(units, 5)
	units = append(units, 6)
	units = append(units, 7)
	units = append(units, 5)
	units = append(units, 6)
	units = append(units, 7)
	units = append(units, 7, 9) // the moment we append 9 capacity increases to 20

	fmt.Println("Slice of units is :", units)
	fmt.Println("Length of unit is :", len(units))
	fmt.Println("Capacity of units is :", cap(units))

	//var stock = []string  we cannot initialize slize like this without any nitial value, we have to pass empty value {}
	var empty_stock = []string{}
	fmt.Println("Capacity of units is :", cap(empty_stock))

	stock := make([]int, 0) //we can use make function to initilaize empty
	fmt.Println("Slice of units is :", stock)
	fmt.Println("Length of unit is :", len(stock))
	fmt.Println("Capacity of units is :", cap(stock))

}
