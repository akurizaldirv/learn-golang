package main

import (
	"errors"
	"fmt"
)

func Pembagian(value int, devider int) (int, error) {
	if devider == 0 {
		return 0, errors.New("Cannot devide by 0")
	} else {
		return value / devider, nil
	}
}

func main() {
	fmt.Println(Pembagian(10, 0))
	fmt.Println(Pembagian(10, 2))
}
