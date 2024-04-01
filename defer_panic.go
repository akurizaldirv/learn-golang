package main

import "fmt"

func logging() {
	fmt.Println("Application Stopped")
}

func runProcess() {
	defer logging()

	fmt.Println("Running Application")
	panic("Got This Error")
	fmt.Println("This should be not executed")
}

func main() {
	runProcess()
}
