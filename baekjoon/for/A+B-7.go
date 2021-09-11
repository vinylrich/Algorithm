//11021번

package main

import "fmt"

func main() {
	var A, B, T int

	var ans []int

	ans = make([]int, T)
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&A, &B)
		ans = append(ans, A+B)
	}
	for i := 0; i < T; i++ {
		fmt.Print("Case #", i+1, ": ", ans[i], "\n")
	}
}
