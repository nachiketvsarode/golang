// in golang, map is a datastructure that provides an unorderd collection of key-value pairs,
//where each key must be unique.
//Maps are used to efficient retrivals of values based on those keys
//Maps stores in key value pairs & in a unorderd way while retriving too it will retun in unordered way only.
//key is index here
//key is most important part of this maps

package main

import "fmt"

func main() {

	// name <---> grades

	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 100
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks of Bob: ", studentGrades["Bob"])
	studentGrades["Bob"] = 100
	fmt.Println("new Marks of bob: ", studentGrades["Bob"])

	//delete(studentGrades, "BOB")
	fmt.Println("Marks of Bob:", studentGrades["Bob"])

	//Checking if a key exists ....exists is a built in feat.
	grades, exists := studentGrades["David"]
	fmt.Println("Grades of David:", grades)
	fmt.Println("David Exists: ", exists)

	Grades, Exists := studentGrades["Prince"]
	fmt.Println("Grades of Prince:", Grades)
	fmt.Println("Prince Exists: ", Exists)

	for index, value := range studentGrades {

		fmt.Printf("Key is %s and value is %d\n", index, value)
	}

	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}

	for index, value := range person {

		fmt.Printf("------------------Key is %s and value is %d\n", index, value)
	}

}
