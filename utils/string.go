package utils

import (
	"strconv"
	"strings"
)

func GetID(id string) uint64 {
	str := strings.Replace(id, "#", "", -1)
	number, err := strconv.ParseUint(str, 0, 64)

	if err != nil {
		return 0
	}

	return number
}
