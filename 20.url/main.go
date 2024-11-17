package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Leanirng URL")
	myurl := "https://eample.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Tyep of URL: %T\n", myurl)

	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Cant parse URL")
		return
	}
	fmt.Printf("Type of URL : %T\n", parsedURL)

	fmt.Println("Host :", parsedURL.Host)
	fmt.Println("Path :", parsedURL.Path)
	fmt.Println("RawQuery :", parsedURL.RawQuery)

	//Modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=nachiketvsarode"

	//constructung a url string from a URL object

	newURL := parsedURL.String()

	fmt.Println("new url is :", newURL)

}
