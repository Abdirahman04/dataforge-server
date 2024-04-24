package services

import "github.com/Abdirahman04/dataforge-server/internal/app/models"

func CreateList(b models.Blueprint) []map[string]interface{} {
  var objs []map[string]interface{}

  for i := 0;i < b.Volume;i++ {
    obj := make(map[string]interface{})
    for _, e := range b.Props {
      var prop models.ProcessedProp
      if e.Type == "string" {
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
