package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	name := NewMap("")

	if name == nil {
		fmt.Println("data kosong:", name)
	} else {
		fmt.Println(name)
	}

}
