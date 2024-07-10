package main

import (
	"application/internal"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	internal.App.Run()
}
