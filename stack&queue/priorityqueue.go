package main

import (
	"fmt"
	"sort"
)

type Queue struct {
	value []int
}

func (q *Queue) Push(add int) {
	q.value = append(q.value, add)
}
func (q *Queue) Pop() int {

	tmp := q.value[0]
	q.value = q.value[1:]
	return tmp
}
func main() {
	priorities := []int{2, 1, 3, 2}
	location := 2
	idx := solution(priorities, location)
	fmt.Println(idx)
}

//location을 어떻게 사용할까?
func solution(priorities []int, location int) int {
	var q Queue
	var cnt int

	var abc []string

	
	q.value = priorities
	a := &q.value[location]//3의 주솟값

	sort.Slice(q.value, func(i, j int) bool {
		return q.value[i] > q.value[j]
	})
	// test := priorities[location]
	// var cnt int = 0
	for i := 0; i < len(priorities); i++ { //오름차순 정렬 후 pop까지 하려면 최대 priorities * qvalue length까지 돌아야함
		// var biggest int
		for j := 1; j < len(q.value); j++ { //0과 1~3까지
			if q.value[0] < q.value[j] { //1 2 3 4
				idx1 := q.Pop() //q.value[0]번의 로케이션 값을 넣음 0x0010
				q.Push(idx1)    //0x0010
				break
			}

		}
	// }

	for i := 0; i < len(priorities); i++ {
		cnt++
		if a == &q.value[i] {
			return cnt
		}
		fmt.Println(q, cnt)
		q.Pop()
		fmt.Println(q, cnt)
	}

	return cnt
}
