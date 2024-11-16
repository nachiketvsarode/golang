package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, orange, banana, prince"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	data1 := "apple. orange. banana. prince"
	parts1 := strings.Split(data1, ".")
	fmt.Println(parts1)

	str := "one two three four five two five two"
	count := strings.Count(str, "two")
	fmt.Println("count", count)

	str = "   Hello, Go !!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed:", trimmed)

	str1 := "prince"
	str2 := "agarwal"
	result := strings.Join([]string{str1, "Kumar", str2}, ", ")
	fmt.Println("results:", result)

}
