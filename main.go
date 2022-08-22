package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 10; i++ {
		if (i % 2) == 0 {
			fmt.Println(`Bilangan ` + strconv.Itoa(i) + ` adalah genap`)
		} else {
			fmt.Println(`Bilangan ` + strconv.Itoa(i) + ` adalah ganjil`)
		}
	}
}
