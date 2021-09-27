package main

import "fmt"

func main() {
	absolutes := []int{4, 7, 12}
	signs := []bool{true, false, true}
	a := solution(absolutes, signs)
	fmt.Println(a)
}
func solution(absolutes []int, signs []bool) int {
	var sum int
	for i, p := range signs {
		if p == true {
			sum += absolutes[i]
		} else {
			sum -= absolutes[i]
		}
	}
	return sum
}
