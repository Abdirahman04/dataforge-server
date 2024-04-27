package inter

import "math/rand"

func RandomIntFromRange(start, end int) int {
  var num int

  for num < start {
    num = rand.Intn(end)
  }

  return num
}
