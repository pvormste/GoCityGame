package util

type Config struct {
	Port    string
	DB_URL  string
	DB_Name string
	Tmpl    string
}

var Conf Config

func init() {
	// Configuration
	Conf = Config{
		Port:    ":3000",
		DB_URL:  "localhost:27017",
		DB_Name: "citygame",
		Tmpl:    "dev",
	}
}
