package main

import (
	// "fmt"
	"fmt"
	"os"

	"github.com/aulianurhady/training/mygram/lib"
	"github.com/aulianurhady/training/mygram/routes"
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
