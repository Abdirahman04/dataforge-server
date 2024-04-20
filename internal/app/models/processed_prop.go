package models

type ProcessedProp struct {
  Id int
  Name string
  Value interface{}
}

func (p ProcessedProp) Addr() (position int, classes []int) {
  return
}
