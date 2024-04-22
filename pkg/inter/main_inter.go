package inter

import (
	"math"
	"math/rand"
)

func RandomNumber(ln int) int {
  return rand.Intn(int(math.Pow(10, float64(ln))))
}
