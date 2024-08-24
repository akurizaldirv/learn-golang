package main

import (
	"fmt"
	"strconv"
	"strings"
)

func polyndrome(text string) string {
	var builder strings.Builder

	for i := len(text) - 1; i >= 0; i-- {
		builder.WriteByte(text[i])
	}

	return builder.String()
}

func main() {
	number := 6354
	converted := strconv.Itoa(number)

	fmt.Println(polyndrome(converted))
	fmt.Println(polyndrome("aku"))
}
