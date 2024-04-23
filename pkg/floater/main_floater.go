package floater

import (
	"math"
	"math/rand"
)

func RandomFloat(ln, decimal int) float64 {
  den := math.Pow(10, float64(decimal))

  num := ln + decimal
  float := rand.Intn(num)

  return float64(float) / den
}
