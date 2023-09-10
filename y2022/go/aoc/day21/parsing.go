package day21

import (
	"strconv"
	"strings"
)

var testinput = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

type Constant int

type Calculation struct {
	lhs, rhs  string
	operation Operation
}

type Operation int

const (
	Add Operation = iota
	Sub
	Mul
	Div
)

func (o Operation) Eval(lhs, rhs int) int {
	switch o {
	case Add:
		return lhs + rhs
	case Sub:
		return lhs - rhs
	case Mul:
		return lhs * rhs
	case Div:
		return lhs / rhs
	}
	panic("unreachable")
}

func OperationFrom(s string) Operation {
	switch s {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mul
	case "/":
		return Div
	}
	panic("unexpected operation: '" + s + "'")
}

func ParseLine(s string) (string, any) {
	// drzm: hmdt - zczc
	splitted := strings.Split(s, ": ")
	lhs := splitted[0]
	rhs := strings.Split(splitted[1], " ")
	if len(rhs) == 1 {
		v, err := strconv.Atoi(rhs[0])
		if err != nil {
			panic(err.Error())
		}
		return lhs, Constant(v)
	}
	if len(rhs) == 3 {
		return lhs,
			Calculation{
				lhs:       rhs[0],
				rhs:       rhs[2],
				operation: OperationFrom(rhs[1]),
			}
	}
	panic("unexpected line: '" + s + "'")
}

func Parse(s string) map[string]any {
	ret := map[string]any{}
	for _, l := range strings.Split(s, "\n") {
		name, expr := ParseLine(l)
		ret[name] = expr
	}
	return ret
}
