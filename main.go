package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// fmt.Print("Hello")
	// fmt.Print("Hello")
	// fmt.Print("Hello")

	// var text1 string = "Hello World"
	// var text2 = "Halo Dunia"
	// text3 := "Hola, Amigos"
	// text4 := text3 + text2

	// fmt.Println(text1)
	// fmt.Println(text2)
	// fmt.Println(text3)

	// var num1 int = 20000
	// var num2 int8 = 127
	// var num20 int8 = 127
	// var num3 = 19000
	// num4 := 20000000

	// fmt.Println(num1 + num3)
	// fmt.Println(num20 + num2) // result : -2 ---cause its int8 type
	// fmt.Println(num4 / 20)
	// fmt.Println(text4)

	// name := "zaldi"
	// age := 25

	// fmt.Printf("Hello, my name is %v and my age is %v \n", name, age)
	// fmt.Printf("Hello, my name is %q and my age is %q \n", name, age)
	// fmt.Printf("Hello, my name is %T and my age is %T \n", name, age)
	// fmt.Printf("Hello, my name is %q and my age is %0.2f \n", name, 25.7897)

	// str := fmt.Sprintf("Hello, my name is %T and my age is %T \n", name, age) // save formatting string
	// fmt.Print(str)

	// Arrays
	var ages [3]int = [3]int{22, 23, 24}
	var ages2 = [3]int{32, 33, 34}
	ages3 := [3]int{42, 43, 44}

	fmt.Println(ages, len(ages))
	fmt.Println(ages2, len(ages2))
	fmt.Println(ages3, len(ages3))

	// Slices : similar to array, but return new slice when append/pop
	var names = []string{"mario", "luigi", "boni"}
	names = append(names, "zaldi")
	fmt.Println(names)

	// slices range
	hobbies := []string{"Fishing", "Running", "Fighting", "Killing"}
	fmt.Println(hobbies[2:])
	fmt.Println(hobbies[:3])
	fmt.Println(hobbies[2:3])

}
