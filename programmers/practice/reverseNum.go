package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution(12345))
}
func solution(n int64) []int {

	s := strconv.Itoa(int(n))
	sArray := strings.Split(s, "")
	var revArr = make([]int, len(sArray))
	var cnt int
	for i := len(sArray) - 1; i >= 0; i-- {
		revArr[i], _ = strconv.Atoi(sArray[cnt])
		cnt++
	}
	return revArr
}
