package main

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
	}
	
}

func main() {
	var l LinkedList
	l.addNode(5)
}
