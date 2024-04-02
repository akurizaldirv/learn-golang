package main

import (
	"encoding/base64"
	"fmt"
)

type Animal struct {
	Name string
	Legs int
}

func (a Animal) toString() string {
	str := fmt.Sprintf("Animal {Name: %v, Legs: %v}", a.Name, a.Legs)
	return str
}

// any = interface{} ----- merupakan interface parent dari semua tipe data
func encrypt(data any) any {
	result := base64.StdEncoding.EncodeToString([]byte(data.(string)))
	return result
}

func decrypt(str string) any {
	result, err := base64.StdEncoding.DecodeString(str)
	if err == nil {
		return string(result)
	}
	fmt.Println(err)
	return str
}

func main() {
	cow := Animal{Name: "Billy", Legs: 4}
	encrypted := encrypt(cow.toString())
	decrypted := decrypt(encrypted.(string))
	fmt.Println(encrypted)
	fmt.Println(decrypted)
}
