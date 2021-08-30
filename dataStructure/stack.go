package main

import "fmt"

type Stack struct {
	value []int
}

func (s *Stack) Push(num int) {
	s.value = append(s.value, num)
}
func (s *Stack) Pop() int {
	idx := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]

	return idx
}
func main() {
	var s Stack

	s.Push(5)
	s.Push(7)
	s.Push(10)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
