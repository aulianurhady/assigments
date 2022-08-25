package main

import (
	"fmt"
	"strconv"
)

func main() {
	friendNames := []string{"Hilmi", "Yosef", "Edi", "Yudha", "Eka", "Thalia", "Sigit", "Cecep", "Rijal", "Bayu"}

	listOfFriends(friendNames...)
}

func listOfFriends(friends ...string) {
	for i, friend := range friends {
		fmt.Println(`Nama ke-` + strconv.Itoa(i) + ` = ` + friend)
	}
}
