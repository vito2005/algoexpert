package main

import (
	"fmt"
)

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return productSum(array, 1)
}
func productSum(array []interface{}, depth int) int {
	result := 0
	for _, v := range array {
		switch v := v.(type) {
		case int:
			result += v
		case SpecialArray:
			result += productSum(v, depth+1)
		}
	}
	return result * depth
}

func main() {

	input := SpecialArray{
		SpecialArray{
			SpecialArray{
				SpecialArray{
					SpecialArray{
						5,
					},
				},
			},
		},
	}
	output := ProductSum(input)
	// expected := 12
	fmt.Println(output)
}
