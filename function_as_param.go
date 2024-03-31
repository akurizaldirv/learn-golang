package main

import (
	"fmt"
	"strings"
)

func main() {
	filter1 := filterDot

	helloFilter("anjing", filter1)
	helloFilter("tai", filterAsterisk)
}

func helloFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func filterDot(name string) string {
	blacklist := []string{"anjing", "tai"}

	if isContain(blacklist, name) {
		return "..."
	} else {
		return name
	}
}

func filterAsterisk(name string) string {
	blacklist := []string{"anjing", "tai"}

	if isContain(blacklist, name) {
		return "***"
	} else {
		return name
	}
}

func isContain(slice []string, str string) bool {
	for _, value := range slice {
		if strings.EqualFold(value, str) {
			return true
		}
	}

	return false
}
