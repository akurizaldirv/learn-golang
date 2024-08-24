package main

import (
	"fmt"
	"strings"
)

func caesar(text string, m int) string {
	var builder strings.Builder

	for _, v := range text {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			offset := int('a') - 1
			if int(v) < int('a') {
				offset = int('A') - 1
			}
			shifted := rune(offset + shift(int(v), m, offset))
			builder.WriteRune(shifted)
		} else {
			builder.WriteString(string(v))
		}
	}
	return builder.String()
}

func getPositiveIndex(number int) int {
	return (26 + (number % 26)) % 26
}

func shift(number int, m int, offset int) int {
	index := ((number + getPositiveIndex(m)) % offset) % 26
	if index == 0 {
		index = 26
	}
	return index
}

func encode(text string, move int) string {
	return caesar(text, move)
}

func decode(text string, move int) string {
	return caesar(text, getInverseMove(move))
}

func getInverseMove(move int) int {
	return (move * -1)
}

func main() {
	text := "ZzAa"
	move := 27
	encrypted := encode(text, move)
	decrypted := decode(encrypted, move)

	fmt.Println(encrypted)
	fmt.Println(decrypted)
}
