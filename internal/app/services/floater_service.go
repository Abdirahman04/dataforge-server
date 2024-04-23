package services

import (
	"strconv"

	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/pkg/floater"
)

func FloaterForge(prop models.Prop) models.ProcessedProp {
  den, _ := strconv.Atoi(prop.Class)

  return models.NewProcessedProp(prop.Id, prop.Name, floater.RandomFloat(prop.Length, den))
}
