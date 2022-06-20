package main

import (
	"fmt"
	"sort"
)

func main() {
	answers := []int{1, 2, 3, 4, 5}
	rst := solution(answers)
	fmt.Println(rst)
}

func bestCandi(answers []int) []int {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 1, 2, 3, 2, 4, 2, 5}
	c := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
	var rst []int
	var acnt, bcnt, ccnt int
	var arst, brst, crst int
	for _, p := range answers {
		if acnt%len(a) == 0 { //4
			acnt = 0
		}
		if bcnt%len(b) == 0 {
			bcnt = 0
		}
		if ccnt%len(c) == 0 {
			ccnt = 0
		}
		if a[acnt] == p {
			arst++
		}
		if b[bcnt] == p {
			brst++
		}
		if c[ccnt] == p {
			crst++
		}
		acnt++
		bcnt++
		ccnt++
	}
	rst = max(arst, brst, crst)
	return rst
}
func max(a, b, c int) []int {
	var rst []int
	var arr []int
	var sorted []int
	arr = append(arr, a, b, c)
	sorted = append(sorted, a, b, c)
	sort.Sort(sort.IntSlice(sorted)) // int 슬라이스를 오름차순으로 정렬
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == sorted[2] {
			rst = append(rst, i+1)
		}
	}
	return rst
}
func solution(answers []int) []int {
	rst := bestCandi(answers)

	return rst
}
