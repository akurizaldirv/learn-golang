package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Application Stopped")
	message := recover()
	fmt.Println("Panic Message : ", message)
}

func runProcess() {
	defer logging()

	fmt.Println("Running Application")
	panic("Got This Error")
}

func main() {
	runProcess()
}
