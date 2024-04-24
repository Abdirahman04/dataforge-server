package services

import (
	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/pkg/booler"
)

func BoolerForge(prop models.Prop) models.ProcessedProp {
  return models.NewProcessedProp(prop.Id, prop.Name, booler.RandomBool())
}
