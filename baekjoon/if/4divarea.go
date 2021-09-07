package main

import "fmt"

func solution(x, y int) int {
	if x < 0 {
		if y < 0 {
			return 3
		}
		return 2
	} else {
		if y < 0 {
			return 4
		}
		return 1
	}
}
func main() {
	x, y := 0, 0
	fmt.Scanf("%d\n%d", &x, &y)
	a := solution(x, y)
	fmt.Println(a)
}
