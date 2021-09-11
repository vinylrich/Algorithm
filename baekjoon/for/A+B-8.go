//11021번

package main

import "fmt"

func main() {
	var T int

	var tmp1, tmp2 int
	var A, B, ans []int

	fmt.Scanln(&T)

	ans = make([]int, T)
	A, B = make([]int, T), make([]int, T)

	for i := 0; i < T; i++ {
		fmt.Scanln(&tmp1, &tmp2)
		A[i] = tmp1
		B[i] = tmp2
		ans[i] = A[i] + B[i]
	}
	for i := 0; i < T; i++ {
		fmt.Print("Case #", i+1, ": ", A[i], " + ", B[i], " = ", ans[i], "\n")
	}
}
