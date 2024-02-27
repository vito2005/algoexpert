package main

import (
	"fmt"
)

// O(n) time | O(1) space
func GetNthFib(n int) int {
	a := 0
	b := 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {

	fmt.Println(GetNthFib(6))
	// 5 // 0, 1, 1, 2, 3, 5
}
