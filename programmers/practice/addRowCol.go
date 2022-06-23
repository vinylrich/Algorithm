package main

import "fmt"

func main() {
	fmt.Println(solution([][]int{{1, 2}, {2, 3}}, [][]int{{3, 4}, {5, 6}}))
}
func solution(arr1 [][]int, arr2 [][]int) [][]int {
	var ans = make([][]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		ans[i] = make([]int, len(arr1[i]))
		for j := 0; j < len(arr1[i]); j++ {
			ans[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	return ans
}
