package dater

import "time"

func RangomDateFromRange(start, end time.Time) time.Time {
  var dt time.Time

  for dt.After(end) || start.After(dt) {
    dt = RandomDate()
  }

  return dt
}
