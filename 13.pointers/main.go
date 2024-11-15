package main

import "fmt"

func modifyValuebyRefrence(num *int) {
	*num = *num * 20 // *num x 20 (multiply)s
}

func main() {

	//first way of declaring
	/*
		var num int
		num = 2
		var ptr *int

		ptr = &num
	*/

	//second way of declaring

	num := 2
	ptr := &num

	fmt.Println("Num has value:", num)
	fmt.Println("pointer contains: ", ptr)

	//to print the value of pointer

	fmt.Println("Data contains through the pointer is :", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyValuebyRefrence(&value)
	fmt.Println("Value contains :", value)

}
