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

func GetConsonants(ln int) string {
  letters := [21]byte{'b','c','d','f','g','h','j','k','l','m','n','p','q','r','s','t','v','w','x','y','z'}

  var str []byte

  for i := 0;i < ln;i++ {
    str = append(str, letters[rand.Intn(26)])
  }

  return string(str)
}

func GetVowels(ln int) string {
  letters := [5]byte{'a','e','i','o','u'}

  var str []byte

  for i := 0;i < ln;i++ {
    str = append(str, letters[rand.Intn(26)])
  }

  return string(str)
}
