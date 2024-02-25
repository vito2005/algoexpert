package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	current := linkedList
	for current.Next != nil {
		if current.Next.Value == current.Value {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return linkedList
}

func main() {
	input := addMany(&LinkedList{Value: 1}, []int{1, 3, 4, 4, 4, 5, 6, 6})
	expected := addMany(&LinkedList{Value: 1}, []int{3, 4, 5, 6})
	actual := RemoveDuplicatesFromLinkedList(input)
	//require.Equal(actual)
	fmt.Println(getValues(actual), expected)
}

func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
	return linkedList
}

func getValues(linkedList *LinkedList) []int {
	values := []int{}
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}
