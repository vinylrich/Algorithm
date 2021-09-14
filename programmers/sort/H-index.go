package main

import (
	"fmt"
	"sort"
)

func main() {
	var citations = []int{3, 0, 6, 1, 5}
	a := solution(citations)
	fmt.Println(a)
}

//6 5 3 1 0
func solution(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	
	if citations[0] == 0 {
		return 0
	}

	for i := 0; i < len(citations); i++ {
		if citations[i] <= i {
			return i
		}
	}
	return len(citations)
}

//알파벳
//은행
// deposit withdrawal loan
