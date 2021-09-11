package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, A int
	var ans []int

	ans = make([]int, N)

	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		ans = append(ans, A)
	}

	sort.Sort(sort.IntSlice(ans))
	fmt.Println(ans[0], ans[len(ans)-1])
}
