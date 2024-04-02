package main

import (
	"errors"
	"fmt"
)

type MathError struct {
	Message string
}

func (error *MathError) Error() string {
	return error.Message
}

func Pembagian(value int, devider int) (int, error) {
	if devider == 0 {
		return 0, errors.New("Cannot devide by 0")
	} else {
		return value / devider, nil
	}
}

func SaveData(id string) error {
	if id == "" {
		return &MathError{Message: "Operation Error"}
	}

	return nil
}

func main() {
	fmt.Println(Pembagian(10, 0))
	fmt.Println(Pembagian(10, 2))

	fmt.Println(SaveData(""))
	fmt.Println(SaveData("1"))

	error := SaveData("")

	switch value := error.(type) {
	case *MathError:
		fmt.Println("math Error", value)
	default:
		fmt.Println("Not Error", value)
	}
}
