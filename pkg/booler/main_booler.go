package booler

import "math/rand"

func RandomBool() bool {
  var bools = [2]bool{true, false}

  return bools[rand.Intn(2)]
}
