package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	breakpoint := 22
	rand.Seed(time.Now().UTC().UnixNano())
	fase := 1
	for {
		isEnded := ambilKorek("Player 1", breakpoint)

		if isEnded {
			break
		}

		fase++
	}
}

func ambilKorek(player string, breakpoint int) (isEnd bool) {
	time.Sleep(1 * time.Second)
	min := 1
	max := 100

	nilaiKorek := rand.Intn(max-min+1) + min

	fmt.Printf("%s mendapat nilai %d \n", player, nilaiKorek)
	result := nilaiKorek % breakpoint

	if result == 0 {
		isEnd = true
		return isEnd
	}

	isEnd = false

	return isEnd
}
