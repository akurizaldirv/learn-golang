package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func (address *Address) ChangeCityToSurabaya() {
	address.City = "Surabaya"
}

func main() {
	address := Address{"Jakarta", "DKI Jakarta", ""}
	fmt.Println(address)

	ChangeCountryToIndonesia(&address)
	address.ChangeCityToSurabaya()
	fmt.Println(address)
}
