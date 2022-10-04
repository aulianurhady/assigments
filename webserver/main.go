package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aulianurhady/training/webserver/service"

	"github.com/gorilla/mux"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)

	r := mux.NewRouter()
	r.HandleFunc("/register", userSvc.RegisterHandler)
	r.HandleFunc("/user", userSvc.GetUserHandler)
	r.HandleFunc("/user/{id}", userSvc.GetUserHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server running on localhost:8080")
	_ = srv.ListenAndServe()
}
