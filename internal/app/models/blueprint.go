package models

import (
	"strconv"
	"time"
)

type Range struct {
  Min string `json:"min"`
  Max string `json:"max"`
}

type Prop struct {
  Name string `json:"name"`
  Type string `json:"type"`
  Length int `json:"length"`
  In []string `json:"in"`
  Range []string `json:"range"`
  Unique bool `json:"unique"`
}

func (p Prop) GetInInt() (arr []int) {
  for _, e := range p.In {
    num, _ := strconv.Atoi(e)
    arr = append(arr, num)
  }

  return
}

func (p Prop) GetInFloat() (arr []float64) {
  for _, e := range p.In {
    num, _ := strconv.ParseFloat(e, 64)
    arr = append(arr, num)
  }

  return
}

func (p Prop) GetInDate() (arr []time.Time) {
  for _, e := range p.In {
    layout := "2008-01-02"
    dateIn, _ := time.Parse(layout, e)
    arr = append(arr, dateIn)
  }

  return
}

type Blueprint struct {
  Volume int `json:"volume"`
  Props []Prop `json:"props"`
}
