package services

import (
	"math/rand"
	"strconv"

	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/pkg/dater"
	"github.com/Abdirahman04/dataforge-server/pkg/inter"
)

func ListForge(b models.Blueprint) []map[string]interface{} {
  var objs []map[string]interface{}



  return objs
}

func CreateList(props []models.Prop) map[string]interface{} {
  newObj := make(map[string]interface{})

  for _, e := range props {
    var prop models.ProcessedProp
    if len(e.In) != 0 {
      prop = GetFromIn(e)
    } else if e.Range.Selected {
      prop = GetFromRange(e)
    } else if e.Type == "string" {
      prop = StringerForge(e)
    } else if e.Type == "int" {
      prop = InterForge(e)
    } else if e.Type == "float" {
      prop = FloaterForge(e)
    } else if e.Type == "bool" {
      prop = BoolerForge(e)
    } else if e.Type == "date" {
      prop = DaterForge(e)
    }
    newObj[prop.Name] = prop.Value
  }

  return newObj
}

func GetFromIn(prop models.Prop) models.ProcessedProp {
  var newProp models.ProcessedProp
  if prop.Type == "string" {
    newProp = models.NewProcessedProp(prop.Id, prop.Name, prop.In[rand.Intn(len(prop.In))])
  } else if prop.Type == "int" {
    var nums []int
    for _, e := range prop.In {
      num, _ := strconv.Atoi(e)
      nums = append(nums, num)
    }
    newProp = models.NewProcessedProp(prop.Id, prop.Name, nums[rand.Intn(len(nums))])
  } else if prop.Type == "float" {
    var nums []float64
    for _, e := range prop.In {
      num, _ := strconv.ParseFloat(e, 64)
      nums = append(nums, num)
    }
    newProp = models.NewProcessedProp(prop.Id, prop.Name, nums[rand.Intn(len(nums))])
  }
  return newProp
}

func GetFromRange(prop models.Prop) models.ProcessedProp {
  var newProp models.ProcessedProp

  if prop.Type == "int" {
    minNum, maxNum := prop.Range.IntR()
    num := inter.RandomIntFromRange(minNum, maxNum)
    newProp = models.NewProcessedProp(prop.Id, prop.Name, num)
  } else if prop.Type == "date" {
    start, end := prop.Range.DateR()
    date := dater.RangomDateFromRange(start, end)
    newProp = models.NewProcessedProp(prop.Id, prop.Name, date)
  }

  return newProp
}
