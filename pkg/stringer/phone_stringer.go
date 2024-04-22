package stringer

import (
	"strconv"

	"github.com/Abdirahman04/dataforge-server/pkg/inter"
)

func GetPhoneNumber() string {
  nm := func() string {
    return strconv.Itoa(inter.RandomNumber(3))
  }
  return "+" + nm() + "-" + nm() + "-" + nm() + "-" + nm()
}
