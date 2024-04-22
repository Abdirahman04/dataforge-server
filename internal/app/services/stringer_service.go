package services

import (
	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/pkg/stringer"
)

func StringerForge(prop models.Prop) models.ProcessedProp {
  var str string
  if prop.Class == "random" {
    str = stringer.GetLetters(prop.Length)
  } else if prop.Class == "word" {
    str = stringer.GetSyllables(prop.Length)
  } else if prop.Class == "email" {
    str = stringer.GetEmail(prop.Length)
  }

  newProp := models.NewProcessedProp(prop.Id, prop.Name, str)
  return newProp
}
