package services

import "github.com/Abdirahman04/dataforge-server/internal/app/models"

func CreateList(b models.Blueprint) []map[string]interface{} {
  var objs []map[string]interface{}

  for i := 0;i < b.Volume;i++ {
    obj := make(map[string]interface{})
    for _, e := range b.Props {
      if e.Type == "string" {
        obj[e.Name] = StringerForge(e).Value
      }
    }
    objs = append(objs, obj)
  }

  return objs
}