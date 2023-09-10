package day25

import (
	"bufio"
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/espang/aoc/aoc"
)

var signal []int

type Operation int

const (
	Increase Operation = iota
	Decrease
	Toggle
	Copy
	Jump
	Out
)

func OpFromString(s string) Operation {
	switch s {
	case "inc":
		return Increase
	case "dec":
		return Decrease
	case "tgl":
		return Toggle
	case "cpy":
		return Copy
	case "jnz":
		return Jump
	case "out":
		return Out
	}
	panic("OpFromString(" + s + ")")
}

type Register struct {
	vs [4]int
}

func (r Register) get(pos byte) int {
	return r.vs[pos]
}

func (r *Register) set(pos byte, val int) {
	r.vs[pos] = val
}

func (r *Register) inc(pos byte) {
	r.vs[pos]++
}

func (r *Register) dec(pos byte) {
	r.vs[pos]--
}

func valueOf(vol ValueOrLocation, r Register) int {
	if vol.IsLocation {
		return r.get(vol.Location)
	}
	return vol.Value
}

type ValueOrLocation struct {
	IsLocation bool
	Location   byte
	Value      int
}

func FromString(s string) ValueOrLocation {
	switch s {
	case "a":
		return ValueOrLocation{Location: 0, IsLocation: true}
	case "b":
		return ValueOrLocation{Location: 1, IsLocation: true}
	case "c":
		return ValueOrLocation{Location: 2, IsLocation: true}
	case "d":
		return ValueOrLocation{Location: 3, IsLocation: true}
	default:
		v, err := strconv.Atoi(s)
		if err != nil {
			panic("FromString(" + s + "):" + err.Error())
		}
		return ValueOrLocation{
			Value: v,
		}
	}
}

type Expr struct {
	operation Operation
	x, y      ValueOrLocation
}

func parseLine(line string) (*Expr, error) {
	splitted := strings.Split(line, " ")
	if len(splitted) == 2 {
		return &Expr{
			operation: OpFromString(splitted[0]),
			x:         FromString(splitted[1]),
		}, nil
	}
	if len(splitted) == 3 {
		return &Expr{
			operation: OpFromString(splitted[0]),
			x:         FromString(splitted[1]),
			y:         FromString(splitted[2]),
		}, nil
	}
	return nil, errors.New("unexpected line format")
}

func execute(e *Expr, r *Register, i int, exprs []*Expr) int {
	switch e.operation {
	case Increase:
		r.inc(e.x.Location)
		return i + 1
	case Decrease:
		r.dec(e.x.Location)
		return i + 1
	case Toggle:
		// idx := i + valueOf(e.x, *r)
		idx := i + r.get(e.x.Location)
		if idx < len(exprs) {
			switch exprs[idx].operation {
			case Increase:
				exprs[idx].operation = Decrease
			case Decrease, Toggle:
				exprs[idx].operation = Increase
			case Copy:
				exprs[idx].operation = Jump
			case Jump:
				exprs[idx].operation = Copy
			}
		}
		return i + 1
	case Copy:
		if e.y.IsLocation {
			v := e.x.Value
			if e.x.IsLocation {
				v = r.get(e.x.Location)
			}
			r.set(e.y.Location, v)
		}
		return i + 1
	case Jump:
		v := e.x.Value
		if e.x.IsLocation {
			v = r.get(e.x.Location)
		}
		if v != 0 {
			return i + valueOf(e.y, *r)
		}
		return i + 1
	case Out:
		v := e.x.Value
		if e.x.IsLocation {
			v = r.get(e.x.Location)
		}
		signal = append(signal, v)
		return i + 1
	}

	panic("execute")
}

func solve(a int, limit int, exprs []*Expr) bool {
	r := &Register{[4]int{a, 0, 0, 0}}
	index := 0
	maxIndex := len(exprs)

	signal = signal[:]
	seen := len(signal)
	next := 0
	for index < maxIndex {
		index = execute(exprs[index], r, index, exprs)

		if len(signal) > seen {
			if len(signal) == limit {
				return true
			}

			if signal[len(signal)-1] == next {
				if next == 0 {
					next = 1
				} else {
					next = 0
				}
				seen = len(signal)
			} else {
				return false
			}
		}
	}
	return false
}

func Part1() (string, error) {
	filename := aoc.MustMakeInputAvailable(context.TODO(), 2016, 25)
	content := aoc.LoadFrom(filename)

	s := bufio.NewScanner(strings.NewReader(content))
	var exprs []*Expr
	for s.Scan() {
		line := s.Text()
		expr, err := parseLine(line)
		if err != nil {
			log.Fatalf("parseLine(%s):%v", line, err)
		}
		exprs = append(exprs, expr)
	}

	if err := s.Err(); err != nil {
		log.Fatal("scanner:", err)
	}
	for a := 0; a < 100_000; a++ {
		if solve(a, 1000, exprs) {
			return strconv.Itoa(a), nil
		}
	}

	return "", errors.New("no solution found")
}

func Part2() (string, error) {
	return "no puzzle", nil
}
