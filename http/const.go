package riphttp

import (
	"strconv"

	"github.com/sharpvik/rip"
)

const ripStatusHeader = "Rip-Status"

var ripStatusOkString = strconv.Itoa(rip.StatusOK)
