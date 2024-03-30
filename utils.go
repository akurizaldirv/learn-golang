package main

// func getInitials(name string) (string, string) {
// 	str := strings.Trim(name, " ")
// 	if len(str) == 0 {
// 		return "_", "_"
// 	}

// 	words := strings.Split(strings.Trim(name, " "), " ")
// 	if len(words) < 1 {
// 		return "_", "_"
// 	} else if len(words) < 2 {
// 		return strings.ToUpper(string(words[0][0])), "_"
// 	} else {
// 		return strings.ToUpper(string(words[0][0])), strings.ToUpper(string(words[len(words)-1][0]))
// 	}
// }

// func sayBye(name string) {
// 	fmt.Println("Good Bye,", name)
// }

// func sayHello(name string) {
// 	fmt.Println("Hello,", name)
// }

// func circleArea(radius float32) float32 {
// 	r2 := float32(math.Pow(float64(radius), 2)) // math.Pow return float64, i need to convert it first before operate
// 	return math.Phi * r2
// }

// func handler(names []string, f func(string)) {
// 	for _, v := range names {
// 		f(v)
// 	}
// }
