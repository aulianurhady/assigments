package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	listPerson := []*Person{
		{Nama: "Nama1", Alamat: "Alamat1", Pekerjaan: "Pekerjaan1", Alasan: "Alasan1"},
		{Nama: "Nama2", Alamat: "Alamat2", Pekerjaan: "Pekerjaan2", Alasan: "Alasan2"},
		{Nama: "Nama3", Alamat: "Alamat3", Pekerjaan: "Pekerjaan3", Alasan: "Alasan3"},
	}

	inputArgs := os.Args

	numberInput, err := strconv.Atoi(inputArgs[1])
	if err != nil {
		fmt.Errorf("Err", "The input from arguments is not number!")
		return
	}

	if numberInput > len(listPerson) {
		fmt.Errorf("Err", "Number from arguments input is out of array length")
		return
	}

	if numberInput <= 0 {
		fmt.Errorf("Err", "Number must be greater than 0")
		return
	}

	printFriendData(numberInput-1, listPerson)
}

func printFriendData(indexSelected int, p []*Person) {
	fmt.Println("Nama: ", p[indexSelected].Nama)
	fmt.Println("Alamat: ", p[indexSelected].Alamat)
	fmt.Println("Pekerjaan: ", p[indexSelected].Pekerjaan)
	fmt.Println("Alasan: ", p[indexSelected].Alasan)
}
