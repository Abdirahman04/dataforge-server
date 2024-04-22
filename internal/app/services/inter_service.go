package services

import (
	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/pkg/inter"
)

func InterForge(prop models.Prop) models.ProcessedProp {
  var num int
  if prop.Class == "random" {
    num = inter.RandomNumber(prop.Length)
  }

  return models.NewProcessedProp(prop.Id, prop.Name, num)
}
