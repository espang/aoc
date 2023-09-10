package day20

import (
	"strconv"
	"strings"
)

var testinput = `1
2
-3
3
-2
0
4`

func Parse(input string) []int {
	ret := []int{}
	for _, line := range strings.Split(input, "\n") {
		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err.Error())
		}
		ret = append(ret, v)
	}
	return ret
}
