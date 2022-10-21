package lib

import (
	"math/rand"
	"time"
)

func GetRandomInt() int {
	min := 1
	max := 100

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}
