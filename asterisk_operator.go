package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer ke address 1
	address3 := &address1 // pointer ke address 1

	fmt.Println("Address 1 :", address1)
	fmt.Println("Address 2 :", address2)
	fmt.Println("Address 3 :", address3)

	address2 = &Address{"Lamongan", "Jawa Timur", "Indonesia"} // change address2 to other pointer
	*address3 = Address{"Surabaya", "Jawa Timur", "Indonesia"} // change value that address3's pointer have, address 1 changed

	fmt.Println("Address 1 :", address1)
	fmt.Println("Address 2 :", address2)
	fmt.Println("Address 3 :", address3)

}
