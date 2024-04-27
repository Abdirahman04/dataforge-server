package dater

import (
	"time"

	"github.com/Abdirahman04/dataforge-server/pkg/inter"
)

func CurrentDate() time.Time {
  return time.Now()
}

func RandomDate() time.Time {
  year := inter.RandomIntFromRange(1, time.Now().Year())
  month := time.Month(inter.RandomIntFromRange(1, 13))
  day := inter.RandomIntFromRange(1, 32)
  hour := inter.RandomIntFromRange(1, 24)
  minu := inter.RandomIntFromRange(1, 61)
  sec := inter.RandomIntFromRange(1, 61)

  return time.Date(year, month, day, hour, minu, sec, 0, time.UTC)
}
