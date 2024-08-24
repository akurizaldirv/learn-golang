package main

import (
	"fmt"
	"strings"
	"unicode"
)

func encode(text string) string {
	var builder strings.Builder

	for _, v := range text {
		if unicode.IsDigit(v) {
			offset := int(rune('0')) - 1
			distance := 10
			builder.WriteRune(rune(getRotIndex(offset, int(v), distance)))
		} else if unicode.IsLetter(v) {
			distance := 26
			offset := int('a') - 1
			if unicode.IsUpper(v) {
				offset = int('A') - 1
			}
			builder.WriteRune(rune(getRotIndex(offset, int(v), distance)))
		} else {
			builder.WriteRune(v)
		}
	}

	return builder.String()
}
func encode(text string) string {
	var builder strings.Builder

	for _, v := range text {
		if unicode.IsDigit(v) {
			offset := int(rune('0')) - 1
			distance := 10
			builder.WriteRune(rune(getRotIndex(offset, int(v), distance)))
		} else if unicode.IsLetter(v) {
			distance := 26
			offset := int('a') - 1
			if unicode.IsUpper(v) {
				offset = int('A') - 1
			}
			builder.WriteRune(rune(getRotIndex(offset, int(v), distance)))
		} else {
			builder.WriteRune(v)
		}
	}

	return builder.String()
}

func getRotIndex(offset int, idx int, distance int) int {
	index := (idx - offset + (distance / 2)) % distance
	if index == 0 {
		index = distance
	}
	return index + offset
}

func main() {
	fmt.Println(encode(strings.ToUpper("abcdefghijklmn !@#_+ opqrstuvwxyz")))
	fmt.Println(encode("abcdefghijklm !@#_+ nopqrstuvwxyz"))
	fmt.Println(encode("01234  !@#_+   56789"))
}
