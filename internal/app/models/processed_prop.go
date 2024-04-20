package models

type ProcessedProp struct {
  Id int
  Name string
  Value interface{}
}

func (p ProcessedProp) Addr() (position int, classes []int) {
  point := 100

  if p.Id > point {
    num := p.Id
    for num > point {
      classes = append(classes, num % point)
      num = int(num / point)
    }
    position = num
  } else {
    position = p.Id
  }
  return
}
