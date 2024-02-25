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
	fast := linkedList
	slow := linkedList
	for fast != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
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
