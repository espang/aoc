package day13

import (
	"fmt"
	"sort"
)

func find(slice []interface{}, element interface{}) int {
	for i, e := range slice {
		if rightOrder(e, element) == nil {
			return i
		}
	}
	return -1
}

func Part2(input string) {
	pairs := parse(input)
	lists := []interface{}{}
	for _, p := range pairs {
		lists = append(lists, p.list1)
		lists = append(lists, p.list2)
	}
	dividerOne := []interface{}{
		[]interface{}{2},
	}
	dividerTwo := []interface{}{
		[]interface{}{6},
	}
	lists = append(lists, dividerOne, dividerTwo)
	sort.Slice(lists, func(i, j int) bool {
		ordered := rightOrder(lists[i], lists[j])
		return *ordered
	})
	i1 := find(lists, dividerOne)
	i2 := find(lists, dividerTwo)
	fmt.Println((i1 + 1) * (i2 + 1))
}
