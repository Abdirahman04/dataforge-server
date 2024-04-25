package services

import (
	"time"

	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/pkg/dater"
)

func DaterForge(prop models.Prop) models.ProcessedProp {
  var date time.Time

  if prop.Class == "now" {
    date = dater.CurrentDate()
  }
  return models.NewProcessedProp(prop.Id, prop.Name, date)
}
