package helper

var Application = "Public App"
var version string

func SayHello(name string) string {
	return "Hello, " + name
}

func sayGoodBye(name string) string {
	return "Good Bye, " + name
}

func GetVersion() string {
	return version
}
