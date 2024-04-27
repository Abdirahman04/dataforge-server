package models

type ProcessedProp struct {
  Id Id
  Name string
  Value interface{}
}

func NewProcessedProp(id Id, name string, val interface{}) ProcessedProp {
  return ProcessedProp{
    Id: id,
    Name: name,
    Value: val,
  }
}
