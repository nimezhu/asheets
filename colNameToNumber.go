package asheets

import (
	"strings"
)

func NumberToColName(idx int) string {
	var s = ""
	for i := idx; i > 0; i = i / 26 {
		s = string(64+i%26) + s
	}
	return s
}
func ColNameToNumber(name string) int {
	if len(name) == 0 {
		return -1 //wrong
	}
	n := strings.ToUpper(name)
	sum := 0
	for i := 0; i < len(n); i++ {
		sum *= 26
		sum += int(n[i]) - int('A') + 1
	}
	return sum
}
