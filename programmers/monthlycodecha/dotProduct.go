package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{-3, -1, 0, 2}
	rst := solution(a, b)

	fmt.Println(rst)

}

func solution(a []int, b []int) int {
	var sum int
	for idx, val := range a {
		sum += val * b[idx]
	}

	return sum
}
