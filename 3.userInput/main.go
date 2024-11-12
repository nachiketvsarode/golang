package main

import (
	"bufio"
	"fmt"
	"os"
)

var name string

func main() {
	//fmt.Println("Hey, Wat's your name ?")
	//fmt.Scan(&name)
	//fmt.Println("Hello, Mr.", name)
	fmt.Println("Hey, whats your name ?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, Mr.", name)

}
