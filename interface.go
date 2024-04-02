package main

import "fmt"

type HasName interface { // creating interface
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello,", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func (a Animal) GetName() string {
	return a.Name
}

func (a Animal) Sound() {
	fmt.Println(a.GetName(), ": GUK GUK")
}

func main() {
	person := Person{Name: "Budi"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)

	animal.Sound()
	fmt.Println(animal.GetName())
}
