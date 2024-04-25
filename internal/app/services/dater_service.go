package services

import (
	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/pkg/dater"
)

func DaterForge(prop models.Prop) models.ProcessedProp {
  return models.NewProcessedProp(prop.Id, prop.Name, dater.CurrentDate())
}
