package genrandom

import (
	"math/rand"
	"time"
)

func GenrandInt(scope int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	num := rand.Intn(scope)
	return num
}
