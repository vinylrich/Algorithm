package main

import (
	"fmt"
)

func main() {
	//case1
	// progresses := []int{93, 30, 55}
	// speeds := []int{1, 30, 5}
	//case2
	progresses := []int{95, 90, 99, 99, 80, 99}
	speeds := []int{1, 1, 1, 1, 1, 1}

	a := solution(progresses, speeds)
	fmt.Println(a)
}

//case1:
func solution(progresses []int, speeds []int) []int { //7 4 9
	//95 90 99 99 80 99
	//4
	cnts := make([]int, len(progresses))
	var rst []int
	for i := 0; i < len(progresses); i++ {
		for j := progresses[i]; j < 100; j += speeds[i] { //0번 인덱스 7번 돎
			cnts[i]++
		}
	}

	//4 4 9
	//7 4 9
	//5 10 1 1 20 1
	for i := 0; i < len(cnts); i++ {
		cnt := 1
		if cnts[i] < 0 {
			continue
		}
		for j := i + 1; j < len(cnts); j++ {
			if cnts[i] >= cnts[j] {
				cnts[j] = -1
				cnt++
			} else {
				break
			}
		}
		rst = append(rst, cnt)
	}
	return rst
}

//case 5 10 1 1 20 1
// 5 10 break rst[0]=1  [1]
// 10 1 -> 10 -1 	5 10 -1 1 20 1 cnt++
// 10 1 -> 10 -1	5 10 -1 -1 20 1 rst=2 rst[1,3]
// 20 1 -> 20 1 cnt 2 [1,3,2]

//case 7 4 9
// 7 4 cnt++ rst=[2] 7<9 break
// 9 rst=1 2,1

//앞에 있는 모든 기능이 완성되지 않으면 안됨
