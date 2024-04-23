package floater

import "math"

func RandomFloat(ln, decimal int) float64 {
  den := math.Pow(10, float64(decimal))

  num := float64(ln + decimal)

  return num / den
}
