package main

import "fmt"

/*
2. 배열에서 모든 중복된 숫자와 횟수를 출력하는 프로그램을 작성하시오.
조건: 1 부터 N 까지의 숫자로 이루어진 배열이 있다. 배열엔 중복된 숫자가 나타날 수 있고, N이 무엇인지는 알 수 없다. (루프 사용은 최소로 한다.)  (0<N<50000)
예: [3:5,4:10,5:1,6:10]
*/
func main() {
	arr := []int{1, 3, 1, 5, 6, 2, 3, 1, 4, 5, 6, 7, 8}
	printDeduplicatedNumber(arr)
}

//ol: overlapped
func printDeduplicatedNumber(arr []int) {
	confirmOL := make(map[int]bool) // int와 이미 있는 값을 판별하기 위해 map을 만들었다.
	oLCnt := make(map[int]int)
	for _, p := range arr {
		//confirmOL[idx]가 이미 있으면 oLCnt[idx]++
		if _, ok := confirmOL[p]; ok {
			oLCnt[p] = oLCnt[p] + 1

		} else {
			confirmOL[p] = true
		}
	}
	// key: overlapped number
	// times: overlapped times
	for oLnum, oLtimes := range oLCnt {
		fmt.Printf("[%d:%d]\n", oLnum, oLtimes+1) //중복된 숫자 + 원본 숫자(1개)
	}
}
