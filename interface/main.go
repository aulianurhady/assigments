package main

import (
	"assigment/service"
	"fmt"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	people := []string{"Hilmi", "Yosef", "Edi", "Yudha", "Eka", "Irfan", "Sigit", "Cecep", "Rijal", "Bayu"}
	for _, person := range people {
		resp := userSvc.Register(&service.User{Name: person})
		fmt.Println(resp)
	}

	usersData := userSvc.GetUser()
	fmt.Println("-----------Hasil get user-------------")
	for _, v := range usersData {
		cetakNama(v.Name)
	}
}

func cetakNama(nama string) {
	fmt.Println(nama)
}
