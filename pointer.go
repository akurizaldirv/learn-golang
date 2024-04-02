package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1  // copy value
	address3 := &address2 // pointer ke address 2

	fmt.Println("Address 1 :", address1)
	fmt.Println("Address 2 :", address2)
	fmt.Println("Address 3 :", address3)

	fmt.Println("AFTER MODIFIED VALUE")

	address1.City = "Bandung"
	address2.Province = "Jawa Tengah" // address 3 ikut berubah
	address3.Country = "Pakistan"     // address 2 ikut berubah

	fmt.Println("Address 1 :", address1)
	fmt.Println("Address 2 :", address2)
	fmt.Println("Address 3 :", address3)

}
