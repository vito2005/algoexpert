package main

import (
	"fmt"
	"sort"
)

// time O(n*logn) | space O(1)
func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	time := 0
	for i := 0; i < len(queries); i++ {
		delay := queries[i]
		time += delay * (len(queries) - 1 - i)
	}
	return time
}

func main() {
	queries := []int{3, 2, 1, 2, 6}
	expected := 17
	actual := MinimumWaitingTime(queries)
	fmt.Println(actual == expected)
}
