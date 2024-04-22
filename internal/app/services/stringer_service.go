package services

import (
	"math/rand"

	"github.com/Abdirahman04/dataforge-server/internal/app/models"
)

func StringerForge(prop models.Prop) models.ProcessedProp {
  chars := [26]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}
  var str []byte

  for i := 0;i < prop.Length;i++ {
    str = append(str, chars[rand.Intn(26)])
  }

  newProp := models.ProcessedProp{
    Id: prop.Id,
    Name: prop.Name,
    Value: string(str),
  }
  return newProp
}
