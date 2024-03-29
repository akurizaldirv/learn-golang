package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// fmt.Print("Hello")
	// fmt.Print("Hello")
	// fmt.Print("Hello")

	var text1 string = "Hello World"
	var text2 = "Halo Dunia"
	text3 := "Hola, Amigos"
	text4 := text3 + text2

	fmt.Println(text1)
	fmt.Println(text2)
	fmt.Println(text3)

	var num1 int = 20000
	var num2 int8 = 127
	var num20 int8 = 127
	var num3 = 19000
	num4 := 20000000

	fmt.Println(num1 + num3)
	fmt.Println(num20 + num2) // result : -2 ---cause its int8 type
	fmt.Println(num4 / 20)
	fmt.Println(text4)

	name := "zaldi"
	age := 25

	fmt.Printf("Hello, my name is %v and my age is %v \n", name, age)
	fmt.Printf("Hello, my name is %q and my age is %q \n", name, age)
	fmt.Printf("Hello, my name is %T and my age is %T \n", name, age)
	fmt.Printf("Hello, my name is %q and my age is %0.2f \n", name, 25.7897)

	str := fmt.Sprintf("Hello, my name is %T and my age is %T \n", name, age) // save formatting string
	fmt.Print(str)

}
