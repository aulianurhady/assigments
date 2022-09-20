package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan bool)
	c <- false
	breakpoint := 22
	rand.Seed(time.Now().UTC().UnixNano())
	fase := 1
	for {
		go ambilKorek("Player 1", breakpoint, c)
		// go ambilKorek("Player 2", breakpoint, c)
		// go ambilKorek("Player 3", breakpoint, c)
		// go ambilKorek("Player 4", breakpoint, c)

		if <-c {
			break
		}
		fase++
	}
}

func ambilKorek(player string, breakpoint int, c chan bool) {
	min := 1
	max := 100

	nilaiKorek := rand.Intn(max-min+1) + min

	fmt.Printf("%s mendapat nilai %d", player, nilaiKorek)
	result := nilaiKorek % breakpoint
	if result == 0 {
		c <- true
	}
	// time.Sleep(2 * time.Second)
}
