package models

import "strconv"

type ProcessedProp struct {
  Id string
  Name string
  Value interface{}
}

func (p ProcessedProp) Addr() (position int, classes []int) {
  point, ln := 100, 2

  position, _ = strconv.Atoi(p.Id[:3])
  if len(p.Id) < ln {
    return
  }
  cls, _ := strconv.Atoi(p.Id[3:])

  if cls > point {
    num := cls
    for num > point {
      classes = append(classes, num % point)
      num = int(num / point)
    }
  }
  return
}
