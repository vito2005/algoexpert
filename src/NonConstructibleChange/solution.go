package main

import (
	"fmt"
	"sort"
)

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	var maxChange int

	for _, coin := range coins {
		if coin > maxChange+1 {
			return maxChange + 1
		}
		maxChange += coin
	}
	return maxChange + 1
}

func main() {

	var coins = []int{5, 7, 1, 1, 2, 3, 22}
	fmt.Println(NonConstructibleChange(coins))

	// 1,2,4 => 8
	// 1,2,5 => 4
	// 5, 7, 1, 1, 2, 3, 22 => 20
}
