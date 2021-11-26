package main

import "fmt"

// 1. 정렬되어 있지 않은 값을 가지는 LinkedList가 주어졌을 때 이 리스트에서 중복되는 원소를 제거하는 코드를 작성하라
// 단 버퍼는 사용할 수 없으며 LinkedList는 직접 구현하여야 한다. 구현된 LinkedList 에는 리스트의 사이즈를 확인 할수 있는 기능을 구현하라.

type Value int
type Node struct {
	next *Node
	val  Value
}
type LinkedList struct {
	root *Node
	tail *Node
}

func (ll *LinkedList) AddNode(val Value) { // 3 4 5 6 7 8
	if ll.root == nil {
		ll.root = &Node{val: val} // 3
		ll.tail = ll.root         // root 3 tail 3
		//tail은 root를 참조
		return
	}
	ll.tail.next = &Node{val: val} //
	ll.tail = ll.tail.next         // ll.tail.next
}

func (ll LinkedList) deduplication() LinkedList {
	keys := make(map[Value]bool)
	ue := []Value{}
	node := ll.root
	for node.next != nil {
		// 이미 node.val을 가지고 있는지 check
		if _, ok := keys[node.val]; ok {

		} else {
			keys[node.val] = true
			ue = append(ue, node.val)
		}
		node = node.next
	}
	//last node(tail node) check
	if _, ok := keys[node.val]; ok {

	} else {
		keys[node.val] = true
		ue = append(ue, node.val)
	}

	newnode := LinkedList{}

	for _, p := range ue {
		newnode.AddNode(p)
	}
	return newnode

}
func PrintNodesSize(ll LinkedList) {
	node := ll.root
	var size int
	for node.next != nil {
		size++
		fmt.Println(node.val)
		node = node.next
	}

	//node가 tail인 경우 tail.next는 존재하지 않기 때문에 node.val 한번 더 출력
	fmt.Println(node.val) //last node

	fmt.Printf("size:%d\n", size)
}

func main() {
	var ll LinkedList

	ll.AddNode(3)
	ll.AddNode(4)
	ll.AddNode(4)
	ll.AddNode(4)
	ll.AddNode(5)
	ll.AddNode(5)
	ll.AddNode(5)
	ll.AddNode(5)
	ll.AddNode(5)
	ll.AddNode(5)
	ll.AddNode(6)
	ll.AddNode(6)
	ll.AddNode(7)

	PrintNodesSize(ll)
	ll = ll.deduplication()
	PrintNodesSize(ll)
}
