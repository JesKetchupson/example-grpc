package config

import (
	"io"
	"os"
)

type Configs struct {
	DBType int
	DBUri  string
	Port   string
	Log    io.Writer
}

func Parse() Configs {
	return Configs{
		DBType: 1,         //os.Getenv("DB_TYPE")  example: "postgres" or parse db_uri
		Port:   ":8080",   //os.Getenv("PORT")
		DBUri:  "file_db", //os.Getenv("DB_URI")
		Log:    os.Stdout,
	}

}
