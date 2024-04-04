package env

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filenames ...string) {
	// In Go, the ellipsis (...) in a function parameter list indicates that the function
	// can accept zero or more strings as filenames. If no filenames are provided
	// it will default to loading the .env file in the current directory.
	err := godotenv.Load(filenames...)
	// Here we use the Load function provided by the godotenv package.
	// This function allows to load environment variables from one or more .env type files
	// Those files need to be provided as arguments to this function.
	if err != nil {
		log.Fatal(err)
	}
	// If an error is catched while running .Load the program is terminated and the error
	// is logged to the console.

	log.Println(".env file successfully loaded")
	// If the files are successfully loaded a message confirmation is printed to the console.
}
