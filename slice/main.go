package main

import (
	"fmt"
	"strconv"
)

func main() {
	friendNames := []string{"Hilmi", "Yosef", "Edi", "Yudha", "Eka", "Thalia", "Sigit", "Cecep", "Rijal", "Bayu"}

	for i, friendName := range friendNames {
		fmt.Println(`Nama ke-` + strconv.Itoa(i) + ` = ` + friendName)
	}
}
