package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("Numbers is :", i)
	}
	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			break
		}
	}
	// the range will give two parameters index and value
	numbers := []int{11, 42, 83, 14, 75}
	for index, value := range numbers {
		fmt.Printf("Index of Number is %d, and value is %d\n", index, value)
	}

	data := "Hello World !"
	for index, value := range data {
		fmt.Printf("Index of Data is %d, and value is %c\n", index, value) // %c is for character because here value will be  a character

	}

}
