package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	newStudent := student{
		firstName: "Zaldi",
		lastName:  "Irvana",
		age:       22,
	}

	newStudent.firstName = "Aldi"
	newStudent.Age = 23

	fmt.Println(newStudent.Age)
	fmt.Println(newStudent.firstName)
}
