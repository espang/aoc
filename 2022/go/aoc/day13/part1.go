package day13

import "fmt"

var testInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func boolPtr(v bool) *bool { return &v }

func rightOrder(lefts, rights interface{}) *bool {
	switch l := lefts.(type) {
	case int:
		switch r := rights.(type) {
		case int:
			if l == r {
				return nil
			}
			return boolPtr(l < r)
		case []interface{}:
			return rightOrder([]interface{}{l}, rights)
		default:
			panic("unexpected type rights!")
		}
	case []interface{}:
		switch r := rights.(type) {
		case int:
			return rightOrder(lefts, []interface{}{r})
		case []interface{}:
			for i, left := range l {
				if len(r) <= i {
					return boolPtr(false)
				}
				ordered := rightOrder(left, r[i])
				if ordered != nil {
					return ordered
				}
			}
			if len(r) == len(l) {
				return nil
			}
			return boolPtr(true)
		default:
			panic("unexpected type rights!")
		}
	default:
		panic("unexpected type lefts!")
	}
}

func Part1(input string) {
	pairs := parse(input)
	total := 0
	for i, p := range pairs {
		ordered := rightOrder(p.list1.([]interface{}), p.list2.([]interface{}))
		if ordered != nil && *ordered {
			total += i + 1
		}
	}
	fmt.Println(total)
}
