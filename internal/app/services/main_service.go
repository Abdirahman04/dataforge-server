package services

import (
	"math/rand"

	"github.com/Abdirahman04/dataforge-server/internal/app/models"
)

func CreateList(b models.Blueprint) []map[string]interface{} {
  var objs []map[string]interface{}

  for i := 0;i < b.Volume;i++ {
    obj := make(map[string]interface{})
    for _, e := range b.Props {
      var prop models.ProcessedProp
      if len(e.In) != 0 {
        prop = GetFromIn(e)
      } else if e.Type == "string" {
        prop = StringerForge(e)
      } else if e.Type == "int" {
        prop = InterForge(e)
      } else if e.Type == "float" {
        prop = FloaterForge(e)
      } else if e.Type == "bool" {
        prop = BoolerForge(e)
      }
      obj[e.Name] = prop.Value
    }
    objs = append(objs, obj)
  }

  return objs
}

func GetFromIn(prop models.Prop) models.ProcessedProp {
  var newProp models.ProcessedProp
  if prop.Type == "string" {
    newProp = models.NewProcessedProp(prop.Id, prop.Name, prop.In[rand.Intn(len(prop.In))])
  }
  return newProp
}
