package main

import "fmt"

func solution(year int) int {
	if year%4 == 0 && year%100 != 0 {
		return 1
	}
	if year%400 == 0 {
		return 1
	}
	return 0
}
func main() {
	year := 0
	fmt.Scanf("%d", &year)
	a := solution(year)

	fmt.Println(a)
}
