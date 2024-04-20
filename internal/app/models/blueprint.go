package models

import (
	"strconv"
	"time"
)

type Range struct {
  Min string `json:"min"`
  Max string `json:"max"`
}

func (r Range) IntR() (min,max int) {
  min, _ = strconv.Atoi(r.Min)
  max, _ = strconv.Atoi(r.Max)

  return
}

func (r Range) FloatR() (min,max float64) {
  min, _ = strconv.ParseFloat(r.Min, 64)
  max, _ = strconv.ParseFloat(r.Max, 64)

  return
}

func (r Range) DateR() (min,max time.Time) {
  layout := "2006-02-01"

  min, _ = time.Parse(layout, r.Min)
  max, _ = time.Parse(layout, r.Max)

  return
}

type Prop struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Type string `json:"type"`
  Length int `json:"length"`
  In []string `json:"in"`
  Range Range `json:"range"`
  Unique bool `json:"unique"`
  Prefix string `json:"prefix"`
  Suffix string `json:"suffix"`
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

func (p Prop) GetPrefixInt() (prefix int) {
  prefix, _ = strconv.Atoi(p.Prefix)
  return
}

func (p Prop) GetPrefixFloat() (prefix float64) {
  prefix, _ = strconv.ParseFloat(p.Prefix, 64)
  return
}

func (p Prop) GetSuffixInt() (prefix int) {
  prefix, _ = strconv.Atoi(p.Suffix)
  return
}

func (p Prop) GetSuffixFloat() (prefix float64) {
  prefix, _ = strconv.ParseFloat(p.Suffix, 64)
  return
}

type Blueprint struct {
  Volume int `json:"volume"`
  Props []Prop `json:"props"`
}

func (b Blueprint) IsValidRangeInt() bool {
  for _, e := range b.Props {
    if e.Unique {
      min, max := e.Range.IntR()
      if max - min <= b.Volume {
        return false
      }
    }
  }

  return true
}

func (b Blueprint) IsValidRangeDate() bool {
  for _, e := range b.Props {
    if e.Unique {
      min, max := e.Range.DateR()
      duration := int((max.Sub(min)).Hours() / 24)
      if duration <= b.Volume {
        return false
      }
    }
  }

  return true
}
