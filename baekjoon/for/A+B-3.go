package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		ans[i] = a + b
	}
	for _, p := range ans {
		fmt.Println(p)
	}
}
