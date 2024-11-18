// using https://jsonplaceholder.typicode.com/todos
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Leanirng CRUD")
	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	res, err := http.Get(myurl)
	if err != nil {
		fmt.Println("Error getting:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response: ", res.Status) // if you want to test this then modify the myrl with incorrect values
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("Data:", string(data))

}
