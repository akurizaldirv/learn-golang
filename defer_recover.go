package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Application Stopped")
}

func runProcess() {
	defer logging()

	fmt.Println("Running Application")
	panic("Got This Error")
	message := recover()
	fmt.Println("Panic Message : ", message)
}

func main() {
	runProcess()
}
