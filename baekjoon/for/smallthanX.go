package main

import "fmt"

func main() {
	var N, X int
	fmt.Scanln(&N, &X)

	var A []int
	A = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < N; i++ {
		if X > A[i] {
			fmt.Print(A[i], " ")
		}
	}
}