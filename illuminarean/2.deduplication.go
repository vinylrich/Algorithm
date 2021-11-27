package main

import "fmt"

func main() {
	arr := []int{1, 3, 1, 5, 6, 2, 3, 1, 4, 5, 6, 7, 8}
	printDeduplicatedNumber(arr)
}

func printDeduplicatedNumber(arr []int) {
	keys := make(map[int]bool)
	cnt := make(map[int]int)
	for _, p := range arr {
		if _, ok := keys[p]; ok {
			cnt[p] = cnt[p] + 1

		} else {
			keys[p] = true
		}
	}

	for key, p := range cnt {
		fmt.Printf("overlaped number:%d, overlapedtimes:%d\n", key, p+1) //중복된 숫자 + 원본 숫자(1개)
	}
}
