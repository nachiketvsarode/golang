// using https://jsonplaceholder.typicode.com/todos
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
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

	fmt.Println("Title response:", todo.Title)
	fmt.Println("completed response:", todo.Completed)

}

func performPostRequest() {

	todo := Todo{
		UserID:    23,
		Title:     "Prince Kumar",
		Completed: true,
	}

	// Convert the Todo struct to JSON

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return

	}

	//convert json data to string

	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos"
	//send post request

	res, err := http.Post(myurl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response:", string(data))

}

func performUpdateRequest() {

	todo := Todo{
		UserID:    237890,
		Title:     "Prince Kumar golang",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return

	}

	//convert json data to string

	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request :", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send the request

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending  request :", err)
		return
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response:", string(data))

}
func main() {
	fmt.Println("Leanirng CRUD")
	performGetRequest()
	performPostRequest()
	performUpdateRequest()

}
