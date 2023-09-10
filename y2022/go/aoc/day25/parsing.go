package day25

import "strings"

var testinput = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func Parse(s string) []string {
	return strings.Split(s, "\n")
}
