// using https://jsonplaceholder.typicode.com/todos
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"titile"`
	Completed bool   `json:"completed"`
}

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

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading:", err)
	// 	return
	// }
	// fmt.Println("Data:", string(data))

	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Todo:", todo)

}
