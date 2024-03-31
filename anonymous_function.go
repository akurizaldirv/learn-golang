package main

import "fmt"

type blacklist func(name string) bool

func register(name string, blacklist blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked,", name)
	} else {
		fmt.Println("Welcome,", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	register("aldo", blacklist)
	register("anjing", func(name string) bool {
		return name == "anjing"
	})
}
