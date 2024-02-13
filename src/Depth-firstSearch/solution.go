package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Name     string
	Children []*Node
}

// time O(V+E)| space O(V)
// V - count of vertices
// E - count of edges
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, node := range n.Children {
		array = node.DepthFirstSearch(array)
	}
	return array
}

func main() {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")
	actual := graph.DepthFirstSearch([]string{})
	expected := []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}
	fmt.Println(reflect.DeepEqual(actual, expected))
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}
