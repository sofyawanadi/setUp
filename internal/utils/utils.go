package utils

import (
	"strconv"
)

func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
