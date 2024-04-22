package stringer

import "math/rand"

func GetLetters(ln int) string {
  letters := [26]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

  var str []byte

  for i := 0;i < ln;i++ {
    str = append(str, letters[rand.Intn(26)])
  }

  return string(str)
}
