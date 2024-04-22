package stringer

import (
	"strconv"

	"github.com/Abdirahman04/dataforge-server/pkg/inter"
)

func GetPhoneNumber() string {
  return "+" + strconv.Itoa(inter.RandomNumber(3)) + strconv.Itoa(inter.RandomNumber(3)) + strconv.Itoa(inter.RandomNumber(3))
}
