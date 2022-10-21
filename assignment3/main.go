package main

import (
	"fmt"

	"github.com/aulianurhady/training/assignment3/routes"
)

func main() {
	port := fmt.Sprintf(":%d", 8080)
	r := routes.CreateRouter()
	r.Run(port)
}
