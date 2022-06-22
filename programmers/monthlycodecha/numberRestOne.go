package main

import "fmt"

func main() {
	fmt.Println(solution(10))
}
func solution(n int) int {
	for i := 1; i < n; i++ {
		if n%i == 1 {
			return i
		}
	}
}

// 다른사람 풀이
//func solution(n int) int {
//	var x int = 1
//
//	for n%x != 1 {
//		x++
//	}
//
//	return x
//}
