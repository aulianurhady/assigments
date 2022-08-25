package main

import "fmt"

type Person struct {
	name string
}

func main() {
	printPerson := func(p []*Person) {
		for i, val := range p {
			fmt.Printf("Nama ke-%d: %s\n", i, val.name)
		}
	}

	var ps = []*Person{
		{name: "Hilmi"},
		{name: "Yosef"},
		{name: "Edi"},
		{name: "Yudha"},
		{name: "Eka"},
		{name: "Thalia"},
		{name: "Sigit"},
		{name: "Cecep"},
		{name: "Rijal"},
		{name: "Bayu"},
	}

	printPerson(ps)
}
