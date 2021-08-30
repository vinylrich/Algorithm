package main

import "fmt"

type Queue []int

func (q Queue) Push(num int) Queue {
	q = append(q, num)
	return q
}
func (q Queue) Pop() (Queue, int) {
	return q[1:], q[0]
}
func main() {
	var s Queue
	var idx int
	s = s.Push(3)
	s = s.Push(4)
	s = s.Push(6)
	s = s.Push(7)
	s = s.Push(8)
	s = s.Push(9)
	s = s.Push(10123)

	s, idx = s.Pop()
	fmt.Println(s, idx)
	s, idx = s.Pop()
	fmt.Println(s, idx)
	s, idx = s.Pop()
	fmt.Println(s, idx)
}
