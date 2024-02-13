package main

import (
	"fmt"
)

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}
	left, right := EvaluateExpressionTree(tree.Left), EvaluateExpressionTree(tree.Right)

	switch tree.Value {
	case -1:
		return left + right
	case -2:
		return left - right
	case -3:
		return left / right
	case -4:
		return left * right
	default:
		return tree.Value
	}
}

func main() {

	tree := &BinaryTree{Value: -1}
	tree.Left = &BinaryTree{Value: 2}
	tree.Right = &BinaryTree{Value: -2}
	tree.Right.Left = &BinaryTree{Value: 5}
	tree.Right.Right = &BinaryTree{Value: 1}
	expected := 6
	actual := EvaluateExpressionTree(tree)
	fmt.Println(actual == expected)
}
