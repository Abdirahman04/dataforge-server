package stringer

import (
	"math/rand"
)

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
    str = append(str, letters[rand.Intn(21)])
  }

  return string(str)
}

func GetVowels(ln int) string {
  letters := [5]byte{'a','e','i','o','u'}

  var str []byte

  for i := 0;i < ln;i++ {
    str = append(str, letters[rand.Intn(5)])
  }

  return string(str)
}

func GetDoubleSyllables(ln int) string {
  con := [10]string{"b","d","g","k","l","m","n","p","s","t"}
  var str string

  for i := 0;i < ln;i++ {
    str += con[rand.Intn(10)] + GetVowels(1)
  }

  return str
}

func GetTripleSyllables(ln int) string {
  cc1 := [4]string{"c","g","s","t"}
  cc2 := [2]string{"h","l"}
  var str string

  for i := 0;i < ln;i++ {
    str += cc1[rand.Intn(4)] + cc2[rand.Intn(2)] + GetVowels(1)
  }

  return str
}

func GetSyllables(ln int) string {
  ec := [10]string{"b","d","g","k","l","m","n","p","s","t"}

  if ln == 1 {
    return GetLetters(1)
  } else if ln == 2 {
    return GetDoubleSyllables(1)
  } else if ln == 3 {
    return GetTripleSyllables(1)
  } else {
    if ln % 2 == 0 {
      return GetDoubleSyllables(ln / 2)
    } else {
      return GetDoubleSyllables(ln / 2) + ec[rand.Intn(10)]
    }
  }
}
