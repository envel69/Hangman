package fonction

import (
	"math/rand"
	"time"
)

func Random(n int) int {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}
