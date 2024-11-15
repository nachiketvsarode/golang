package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {

	var prince Person
	fmt.Println("Prince Person :", prince)

	prince.FirstName = "Prince"
	prince.LastName = "Agarwal"
	prince.Age = 32

	//fmt.Println("Prince Person: ", prince)

	person1 := Person{
		FirstName: "Akash",
		LastName:  "Singh",
		Age:       26,
	}
	fmt.Println("First name of  person1 is : ", person1.FirstName)

	var person2 = new(Person)

	person2.FirstName = "Simran"
	person2.LastName = "Agarwal"
	person2.Age = 28

	fmt.Println("First name of  person2 is : ", person2.FirstName)

	var employee1 Employee

	employee1.Person_Details = Person{
		FirstName: "Prince",
		LastName:  "Agarwal",
		Age:       26,
	}

	employee1.Person_Contact.Email = "abcd@gmail.com"
	employee1.Person_Contact.Phone = "12345678"

	employee1.Person_Address = Address{
		House: 12,
		Area:  "Ranchi",
		State: "Jharkhand",
	}

	fmt.Println("Employee 1 : ", employee1.Person_Contact.Email)
	fmt.Println("Employee 1 : ", employee1.Person_Address.State)

}
