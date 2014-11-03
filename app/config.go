package app

// ### INIT

// Stores the configuration and initialize
var Config *AppConfiguration

func init() {
	Config = &AppConfiguration{
		Port:    ":3000",
		DB_URL:  "localhost:27017",
		DB_Name: "citygame",
		Tmpl:    "dev",
		DB:      &DatabaseConfiguration{},
	}
}
