package main

import (
	"fmt"
	"main/database"
	"main/helper"     // sesuai dengan yang di go.mod
	_ "main/internal" // blank identifier (_) : package tidak digunakan, dipanggil hanya untuk memanggil function init
)

func main() {
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // lower case on first letter: private: cannot accessed outside package
	fmt.Println(helper.SayHello("Eko"))
	// fmt.Println(helper.sayGoodBye("Eko")) // lower case on first letter: private: cannot accessed outside package
	fmt.Println(helper.GetVersion())

	fmt.Println(database.GetName())
}
