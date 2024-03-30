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
	// var ages [3]int = [3]int{22, 23, 24}
	// var ages2 = [3]int{32, 33, 34}
	// ages3 := [3]int{42, 43, 44}

	// fmt.Println(ages, len(ages))
	// fmt.Println(ages2, len(ages2))
	// fmt.Println(ages3, len(ages3))

	// // Slices : similar to array, but return new slice when append/pop
	// var names = []string{"mario", "luigi", "boni"}
	// names = append(names, "zaldi")
	// fmt.Println(names)

	// // slices range
	// hobbies := []string{"Fishing ad Shad", "Running", "Fighting", "Killing"}
	// fmt.Println(hobbies[2:])
	// fmt.Println(hobbies[:3])
	// fmt.Println(hobbies[2:3])
	// fmt.Println(hobbies[0])

	// hello := "Hello my friends!"
	// fmt.Println(strings.Contains(hello, "ell"))            // true
	// fmt.Println(strings.Contains(hello, "elo"))            // false
	// fmt.Println(strings.ToUpper(hello))                    // HELLO MY FRIENDS!
	// fmt.Println(strings.ReplaceAll(hello, "my", "fookin")) // Hello fookin friends!
	// fmt.Println(strings.Index(hello, "my"))                // 6
	// fmt.Println(strings.Index(hello, "myass"))             // -1
	// fmt.Println(hello[2])                                  // 108
	// fmt.Println(string(rune(hello[2])))                    // l
	// fmt.Println(hello[2:10])                               // llo my f
	// fmt.Println(hello[:8])                                 // Hello my

	// newAges := []int{20, 12, 13, 44, 23, 19, 29}

	// sort.Ints(newAges)                         // not return anything, but affect newAges
	// fmt.Println(newAges)                       // [12 13 19 20 23 29 44]
	// fmt.Println(sort.SearchInts(newAges, 20))  // 3 (index after sorted)
	// fmt.Println(sort.SearchInts(newAges, 999)) // 7 (when not found, returned length -- which is index out of bounds)

	// sort.Strings(hobbies)                                // not return anything, but affect hobbies
	// fmt.Println(hobbies)                                 // [Fighting Fishing ad Shad Killing Running]
	// fmt.Println(sort.SearchStrings(hobbies, "Running"))  // 3 (index after sorted)
	// fmt.Println(sort.SearchStrings(hobbies, "Runnings")) //when not found, returned length -- which is index out of bounds

	// cities := []string{"Lamongan", "Surabaya", "Malang", "Pati", "Surakarta", "Jakarta", "Bandung"}

	// x := 0
	// for x < 5 {
	// 	fmt.Println("while loop value is x =", x)
	// 	x++
	// }

	// for x := 0; x < 5; x++ {
	// 	fmt.Println("for loop value is x =", x)
	// }

	// for _, v := range cities { // replace index with _ to avoid error from not using index
	// 	fmt.Println(v)
	// }

	// for i, v := range cities { // using index
	// 	fmt.Printf("Index: %v --- Value: %v \n", i, v)
	// }

	// age := 35
	// fmt.Println(age == 45)
	// fmt.Println(age <= 50)
	// fmt.Println(age >= 40)
	// fmt.Println(age != 45)

	// for i, v := range cities {
	// 	if i == 2 {
	// 		fmt.Println("Past this value")
	// 		continue
	// 	} else if i > 3 {
	// 		fmt.Println("Break after this")
	// 		break
	// 	}

	// 	fmt.Printf("Value at index: %v is %v \n", i, v)
	// }

	// names := []string{"Bobi", "Budi", "Buddha", "Binar"}
	// handler(names, sayBye)
	// handler(names, sayHello)
	// // handler(names, circleArea()) // error, doesn't have mismatch type of parameter in parameter
	// fmt.Println(circleArea(3.5))
	// fmt.Println(circleArea(3))

	// fmt.Println(string(names[0][2]))

	// fmt.Println(getInitials("  "))
	// fmt.Println(getInitials("Bobi"))
	// fmt.Println(getInitials("Bobi Ilda"))

	// menu := map[string]float32{
	// 	"salad":    6.44,
	// 	"pear":     2.44,
	// 	"burger":   10.44,
	// 	"sandwich": 7.44,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["pear"])

	// phonebook := map[int]string{
	// 	123123: "Luigi",
	// 	909090: "Police",
	// 	209083: "Bomber",
	// }

	// fmt.Println(phonebook)
	// fmt.Println(phonebook[123123])

	// phonebook[123123] = "Doni"
	// phonebook[909090] = "Boni"
	// phonebook[908978] = "Minnah"

	// fmt.Println(phonebook)

	// var fruits = make([]string, 2)
	// fmt.Println(fruits[0] == "")

	// for key, val := range phonebook {
	// 	fmt.Printf("Key: %v --- Val: %v \n", key, val)
	// }

	name := "zaldi"
	update(name)
	fmt.Println(name)

	fmt.Println(&name)
	updatePointer(&name) // & mean get memory addres/pointer from variable
	fmt.Println(name)
}

func update(name string) {
	name = "new name" // update name, but not returned, so it doesn't affect other scope
}

func updatePointer(name *string) {
	fmt.Println("Pointer :", name)
	*name = "new name by pointer" // update direct to pointer/memory address
	// // * mean get value from a pointer
}
