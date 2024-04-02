package database

var connection string

func init() {
	connection = "MySQL"
}

func GetName() string {
	return connection
}
