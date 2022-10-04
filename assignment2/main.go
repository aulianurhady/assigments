package main

import (
	"fmt"
	"os"

	"github.com/aulianurhady/training/assignment2/lib"

	"github.com/aulianurhady/training/assignment2/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		os.Exit(1)
	}
}

func main() {
	lib.InitDatabase()

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	r := routes.CreateRouter()
	r.Run(port)
}
