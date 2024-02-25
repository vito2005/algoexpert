package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space
func MiddleNode(linkedList *LinkedList) *LinkedList {
	current := linkedList
	middle := linkedList
	length := 0
	for current != nil {
		length++
		current = current.Next
	}

	for i := 0; i < length/2; i++ {
		middle = middle.Next
	}
	return middle
}

func main() {
	linkedList := &LinkedList{Value: 0}
	linkedList.Next = &LinkedList{Value: 1}
	expected := &LinkedList{Value: 2}
	linkedList.Next.Next = expected
	expected.Next = &LinkedList{Value: 3}
	actual := MiddleNode(linkedList)
	fmt.Println(actual)
}
