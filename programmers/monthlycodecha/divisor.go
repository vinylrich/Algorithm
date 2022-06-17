package main

import "fmt"

const odd = 1
const even = 2

func discrim(num int) int {
	var cnt int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			cnt++
		}
	}
	if cnt%2 == 0 {
		return 2
	} else {
		return 1
	}
}
func main() {
	rst := solution(13, 17)
	fmt.Println(rst)
}
func solution(left int, right int) int {
	var sum int
	for i := left; i <= right; i++ {
		rst := discrim(i)
		if rst == odd {
			sum -= i
		}
		if rst == even {
			sum += i
		}
	}
	return sum
}
