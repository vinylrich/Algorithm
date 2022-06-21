package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution(1, 20, 4))
}

func solution(price int, money int, count int) int64 {
	cngPrice := price
	var rst float64
	for i := 0; i < count; i++ {
		money -= cngPrice //20-3
		cngPrice += price
	}
	if money > 0 {
		return 0
	}
	rst = math.Abs(float64(money))
	return int64(rst)
}
