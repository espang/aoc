package day23

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const input = `set b 81
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
sub b -17
jnz 1 -23`

var registerKeys = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

type Register struct {
	m    map[string]*big.Int
	jump int
}

func (r Register) String() string {
	var ss []string
	for k, v := range r.m {
		ss = append(ss, k+":"+v.Text(10))
	}
	return strings.Join(ss, " ")
}

type Operation interface {
	Update(*Register)
}

type SetValue struct {
	Register string
	Value    *big.Int
}

func (s SetValue) Update(reg *Register) {
	reg.m[s.Register] = s.Value
}

type SetPos struct {
	Register string
	Pos      string
}

func (s SetPos) Update(reg *Register) {
	reg.m[s.Register] = reg.m[s.Pos]
}

type SubValue struct {
	Register string
	Value    *big.Int
}

func (s SubValue) Update(reg *Register) {
	i := &big.Int{}
	reg.m[s.Register] = i.Sub(reg.m[s.Register], s.Value)
}

type SubPos struct {
	Register string
	Pos      string
}

func (s SubPos) Update(reg *Register) {
	i := &big.Int{}
	reg.m[s.Register] = i.Sub(reg.m[s.Register], reg.m[s.Pos])
}

type MulValue struct {
	Register string
	Value    *big.Int
}

func (s MulValue) Update(reg *Register) {
	i := &big.Int{}
	reg.m[s.Register] = i.Mul(reg.m[s.Register], s.Value)
}

type MulPos struct {
	Register string
	Pos      string
}

func (s MulPos) Update(reg *Register) {
	i := &big.Int{}
	reg.m[s.Register] = i.Mul(reg.m[s.Register], reg.m[s.Pos])
}

type JumpValue struct {
	Value *big.Int
	Jump  int
}

func (s JumpValue) Update(reg *Register) {
	if s.Value.BitLen() != 0 {
		reg.jump = s.Jump
	}
}

type JumpPos struct {
	Pos  string
	Jump int
}

func (s JumpPos) Update(reg *Register) {
	if reg.m[s.Pos] != nil && reg.m[s.Pos].BitLen() != 0 {
		reg.jump = s.Jump
	}
}

func ParseLine(l string) (Operation, error) {
	parts := strings.Split(l, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("unexpected length of %d with %q", len(parts), l)
	}

	isKey := func(s string) bool {
		for _, v := range registerKeys {
			if s == v {
				return true
			}
		}
		return false
	}
	register := parts[1]
	switch parts[0] {
	case "set":
		if !isKey(register) {
			return nil, fmt.Errorf("unexpected register pos %s in %q", register, l)
		}
		if isKey(parts[2]) {
			return SetPos{Register: register, Pos: parts[2]}, nil
		} else {
			v, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unexpected value %s in %q: %w", parts[2], l, err)
			}
			return SetValue{Register: register, Value: big.NewInt(v)}, nil
		}
	case "sub":
		if !isKey(register) {
			return nil, fmt.Errorf("unexpected register pos %s in %q", register, l)
		}
		if isKey(parts[2]) {
			return SubPos{Register: register, Pos: parts[2]}, nil
		} else {
			v, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unexpected value %s in %q: %w", parts[2], l, err)
			}
			return SubValue{Register: register, Value: big.NewInt(v)}, nil
		}
	case "mul":
		if !isKey(register) {
			return nil, fmt.Errorf("unexpected register pos %s in %q", register, l)
		}
		if isKey(parts[2]) {
			return MulPos{Register: register, Pos: parts[2]}, nil
		} else {
			v, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unexpected value %s in %q: %w", parts[2], l, err)
			}
			return MulValue{Register: register, Value: big.NewInt(v)}, nil
		}
	case "jnz":
		jump, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unexpected value %s in %q: %w", parts[2], l, err)
		}

		if isKey(parts[1]) {
			return JumpPos{Pos: register, Jump: int(jump)}, nil
		} else {
			v, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unexpected value %s in %q: %w", parts[1], l, err)
			}
			return JumpValue{Value: big.NewInt(v), Jump: int(jump)}, nil
		}
	default:
		return nil, fmt.Errorf("unexpected instruction %s in %q", parts[0], l)
	}
}

func Run(reg Register, ops []Operation, idx int) (Register, map[int]int) {
	counter := map[int]int{}
	steps := uint64(0)
	for idx >= 0 && idx < len(ops) {
		if steps%100_000_000 == 0 {
			fmt.Println(steps, reg)
		}
		steps++
		counter[idx]++
		ops[idx].Update(&reg)
		if reg.jump != 0 {
			idx += reg.jump
			reg.jump = 0
		} else {
			idx++
		}
	}
	return reg, counter
}

func part1(ops []Operation) {
	register := map[string]*big.Int{}
	for _, k := range registerKeys {
		register[k] = big.NewInt(0)
	}
	reg, counter := Run(Register{
		m: register,
	}, ops, 0)
	fmt.Println(reg)
	fmt.Println(counter)
}

func part2(ops []Operation) *big.Int {
	// when start at 0
	// b = 81 * 100 - 100000 = -19000
	// c = -17000 - 17000 = -34000
	register := map[string]*big.Int{}
	for _, k := range registerKeys {
		register[k] = big.NewInt(0)
	}
	register["a"] = big.NewInt(1)

	reg, _ := Run(Register{
		m: register,
	}, ops, 0)

	return reg.m["h"]
}
func operations() []Operation {
	lines := strings.Split(input, "\n")
	var operations []Operation

	for i, l := range lines {
		o, err := ParseLine(l)
		if err != nil {
			fmt.Println("error in line ", i)
			fmt.Println(err)
			os.Exit(1)
		}
		operations = append(operations, o)
	}

	return operations
}

func Part1() (string, error) {
	part1(operations())
	return "", nil
}

func Part2() (string, error) {
	counter := 0
	for b := 108100; b <= 125100; b += 17 {
		v := big.NewInt(int64(b))
		if !v.ProbablyPrime(0) {
			counter++
		}
	}
	return strconv.Itoa(counter), nil
}
