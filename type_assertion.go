package main

import (
	"fmt"
	"strconv"
)

func random() any {
	return true
}

func main() {
	// result := random()

	// resultArr := strings.Split(result.(string), "") // if result is Integer, then this will throws error
	// fmt.Println(resultArr)

	// switch value := result.(type) {
	// case string:
	// 	fmt.Println("String", value)
	// case int:
	// 	fmt.Println("Int", value)
	// default:
	// 	fmt.Println("Unknown", value)
	// }

	value := "20"
	// fmt.Println(value.(int)) // this will throws error
	result, _ := strconv.Atoi(value)
	fmt.Println(result)
}
