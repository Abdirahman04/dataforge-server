package services

import "github.com/Abdirahman04/dataforge-server/internal/app/models"

func InterForge(prop models.Prop) models.ProcessedProp {
  var num int
  if prop.Class == "random" {
    num = randomIntLength(prop.Length)
  }

  return models.NewProcessedProp(prop.Id, prop.Name, num)
}
