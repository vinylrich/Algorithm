package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 0}
	a := solution(numbers)
	fmt.Println(a)
}

func sum(numbers []int) int {
	var sum int
	for _, p := range numbers {
		sum += p
	}
	return sum
}
func solution(numbers []int) int {
	arrsum := sum(numbers)
	rst := 45 - arrsum
	return rst
}
