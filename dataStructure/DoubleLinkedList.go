package main

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}
type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) addNode(val int) {
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{Val: val} // Tail 3-> 5
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev
}

func (l *LinkedList) PrintNode() {
	node := l.Root
	for node.Next != nil {
		fmt.Print(node.Val, "->")
		node = node.Next
	}
	fmt.Println(node.Val)
}
func (l *LinkedList) Remove(node *Node) { //3 4 5 6 7
	if node == nil {
		panic("지우려고 하는 node가 없습니다.")
	}
	if node == l.Root {
		l.Root = l.Root.Next
		if l.Root != nil {
			l.Root.Prev = l.Root
		}
		node.Next = nil
		return
	}
	prev := node.Prev

	if node == l.Tail {
		prev.Next = nil
		l.Tail = nil
		l.Tail.Prev = prev
	} else {
		node.Prev = nil
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
	}
	node.Next = nil

}
func main() {
	var l LinkedList
	l.addNode(3)
	l.addNode(6)
	l.addNode(7)

	l.PrintNode()
	l.Remove(l.Root)
	l.PrintNode()
}
