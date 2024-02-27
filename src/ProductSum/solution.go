package main

import (
	"fmt"
)

type SpecialArray []interface{}

type StackObj struct {
	obj   interface{}
	depth int
}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	result := 0
	depth := 0
	stack := []StackObj{{
		obj:   array,
		depth: depth,
	}}

	for len(stack) > 0 {
		i := len(stack) - 1

		last := stack[i]
		stack = stack[:len(stack)-1]

		obj, depth := last.obj, last.depth
		switch v := obj.(type) {
		case int:
			multiplier := getMultiplier(depth)
			result += multiplier * v
		case SpecialArray:
			for _, el := range v {
				newElem := StackObj{obj: el, depth: depth + 1}
				stack = append(stack, newElem)
			}
		case []interface{}:
			for _, el := range v {
				newElem := StackObj{obj: el, depth: depth + 1}
				stack = append(stack, newElem)
			}
		}
	}
	fmt.Println(result)
	return result
}

func getMultiplier(depth int) int {
	multiplier := 1
	for i := 2; i <= depth; i++ {
		multiplier *= i
	}
	return multiplier
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
	expected := 12
	fmt.Println(output == expected)
}
