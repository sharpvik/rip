package proto

import (
	"strconv"
)

func intWithSpaceAsBytes(i int) []byte {
	return []byte(strconv.Itoa(i) + " ")
}
