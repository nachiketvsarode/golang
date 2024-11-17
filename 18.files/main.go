package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt")

		if err != nil {
			fmt.Println("Error while creating file: ", err)
			return
		}
		defer file.Close()
	*/
	/*
		once func main() then close will be run, this function is important in many cases when you are done working with file.
		when you oopen a file using function like os.create, os.Open orthers you are utilizing system reosurces to interact with that file.
		failing to close could leak resources & cause iissues like running out of descriptors.
		the Close() function is responsible for releasing  those resources.
	*/
	/*
		content := "hello world by prince"
		byte, errors := io.WriteString(file, content+"\n")
		fmt.Fprintln(os.Stdout, []any{"byte written:", byte}...)
		if errors != nil {
			fmt.Println("Error while writing file:", errors)
			return
		}

		fmt.Println("successfully created file")

	*/

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()

	//create a buffer to read file content

	buffer := make([]byte, 1024)

	// Read the file content inot the buffer

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error while reading file", err)
			return
		}

		fmt.Println(string(buffer[:n]))
	}

}
