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

func GetSyllables(ln int) string {
  cc1 := [5]byte{'c','g','s','t','n'}
  cc2 := [2]byte{'h','l'}
  ec := [10]byte{'b','d','g','k','l','m','n','p','s','t'}

  if ln == 1 {
    return GetLetters(1)
  } else if ln == 2 {
    return ec[rand.Intn(10)] + GetVowels(1)
  } else if ln == 3 {
    return cc1[rand.Intn(5)] + cc2[rand.Intn(2)] + GetVowels(1)
  } else {
    return GetLetters(ln)
  }
}
