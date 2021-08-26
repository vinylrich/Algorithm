package main

import "fmt"

func main() {
	a := solution(4, 7)
	fmt.Println(a)
}
func solution(a int, b int) int64 {
	var sum int64
	big, small := CngNum(a, b)

	for i := small; i <= big; i++ { //4,7 4+5+6+7 9 15 22
		if small == big {
			return int64(small)
		}
		sum = sum + int64(i) //
	}
	return sum
}

func CngNum(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}
