package service

import (
	"math/rand"
	"time"
)

func SetupDices() {
	rand.Seed(time.Now().Unix())
}

func RollDice(n int) int {

	return rand.Intn(n) + 1
}
