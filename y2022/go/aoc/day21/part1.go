package day21

import (
	"container/list"
	"fmt"
)

func solve(name string, namesExpr map[string]any) int {
	cache := map[string]int{}

	stack := list.New()
	stack.PushFront(name)
	for stack.Len() > 0 {
		current := stack.Front().Value.(string)
		currentExpr := namesExpr[current]
		switch expr := currentExpr.(type) {
		case Constant:
			cache[current] = int(expr)
			stack.Remove(stack.Front())
		case Calculation:
			// do we know the values?
			lhs, ok1 := cache[expr.lhs]
			rhs, ok2 := cache[expr.rhs]
			if ok1 && ok2 {
				cache[current] = expr.operation.Eval(lhs, rhs)
				stack.Remove(stack.Front())
			} else {
				if !ok1 {
					stack.PushFront(expr.lhs)
				}
				if !ok2 {
					stack.PushFront(expr.rhs)
				}
			}
		}
	}
	return cache[name]
}

func Part1(input string) {
	nameToExpr := Parse(input)
	fmt.Println(solve("root", nameToExpr))
}

func Part2(input string) {
	nameToExpr := Parse(input)
	root := nameToExpr["root"].(Calculation)
	delete(nameToExpr, "root")
	delete(nameToExpr, "humn")

	rhs := solve(root.rhs, nameToExpr)
	// experimentation:
	// lhs is greater than rhs for humn:0 and
	// strictly monotonic falling for increasing numbers.
	// binary search between a value for which lhs > rhs
	// and one for which lhs < rhs:
	lower := 0
	upper := 5_000_000_000_000
	for i := 0; i < 100; i++ {
		// There could be edge cases in which this doesn't find the correct solution.
		// But it works here
		mid := ((upper - lower) / 2) + lower

		nameToExpr["humn"] = Constant(mid)
		lhs := solve(root.lhs, nameToExpr)
		if lhs == rhs {
			fmt.Println(mid)
			break
		}
		if lhs > rhs {
			lower = mid
		}
		if lhs < rhs {
			upper = mid
		}
	}
}
