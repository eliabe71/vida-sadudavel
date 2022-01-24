package utils

import (
	"strconv"
	"strings"
)

func BiggerThen(a string, b string) (string, bool) {
	va := strings.Split(a, ":")
	vb := strings.Split(b, ":")
	ha, _ := strconv.Atoi(va[0])
	hb, _ := strconv.Atoi(vb[0])
	if ha == hb {
		ma, _ := strconv.Atoi(va[1])
		mb, _ := strconv.Atoi(vb[1])
		if ma == mb {
			return a, true
		}
		if mb > ma {
			return b, false
		}
		return a, false
	}
	if hb > ha {
		return b, false
	}

	return a, false
}
