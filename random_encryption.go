package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func generateKey() string {
	rand.NewSource(time.Now().UnixNano())
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	key := strings.Split(alphabet, "")

	rand.Shuffle(len(alphabet), func(i, j int) {
		key[i], key[j] = key[j], key[i]
	})

	return strings.Join(key, "")
}

func encode(key string, text string) string {
	var builder strings.Builder

	for _, v := range text {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			isUpper := false
			offset := int('a')
			if v < 'a' {
				isUpper = true
				offset = int('A')
			}

			index := int(v) - offset
			if isUpper {
				builder.WriteString(strings.ToUpper(string(key[index])))
			} else {
				builder.WriteString(string(key[index]))
			}
		} else {
			builder.WriteRune(v)
		}
	}

	return builder.String()
}

func decode(key string, text string) string {
	var builder strings.Builder

	for _, v := range text {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			index := strings.IndexRune(key, unicode.ToLower(v))

			isUpper := false
			offset := int('a')
			if v < 'a' {
				isUpper = true
				offset = int('A')
			}

			decodedChar := rune(index + offset)
			if isUpper {
				builder.WriteString(strings.ToUpper(string(decodedChar)))
			} else {
				builder.WriteString(string(decodedChar))
			}
		} else {
			builder.WriteRune(v)
		}
	}

	return builder.String()
}

func main() {
	key := generateKey()
	encoded := encode(key, "Hello World!")
	decoded := decode(key, encoded)

	fmt.Println(key)
	fmt.Println(encoded)
	fmt.Println(decoded)
}
