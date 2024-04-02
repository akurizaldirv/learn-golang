package main

import "fmt"

type Student struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	student := Student{
		FirstName: "Zaldi",
		LastName:  "Irvana",
		Age:       22,
	}

	fmt.Println(student)
}
