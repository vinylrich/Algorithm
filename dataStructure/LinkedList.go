package main

import "fmt"

type Node struct {
	Next  *Node
	Value int
}
type LinkedList struct {
	Root *Node
	Tail *Node
}

//5 -> 3 -> 4 -> 7
func (ll *LinkedList) AddNode(val int) {
	if ll.Root == nil {
		ll.Root = &Node{Value: val}
		ll.Tail = ll.Root
		return
	}
	ll.Tail.Next = &Node{Value: val}
	ll.Tail = ll.Tail.Next
}
func (ll *LinkedList) Remove(node *Node) {
	if node == nil {
		panic("Empty node")
	}
	if node == ll.Root {
		ll.Root = ll.Root.Next
	}
	//5 3 7 23
	prev := ll.Root
	for prev.Next != node { //23이 들어있는 노드
		prev = prev.Next
	}
	if node == ll.Tail {
		prev.Next = nil
		ll.Tail = prev
	} else {
		prev.Next = prev.Next.Next //3 4 10
	}
}
func (ll *LinkedList) PrintNodes() {
	node := ll.Root

	for node.Next != nil {
		fmt.Printf("%d ->", node.Value)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Value)
}
func main() {
	var a LinkedList
	a.AddNode(5)
	a.AddNode(3)
	a.AddNode(7)
	a.AddNode(23)
	a.PrintNodes()
	a.Remove(a.Tail)

	a.PrintNodes()
}
