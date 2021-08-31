package main

/*
선택정렬의 프로세스
i부터 최솟값을 찾음

최솟값이 나오면 i와 j 스왑
*/
import "fmt"

func main() {
	var num = []int{9, 6, 7, 3, 5}
	a := solution(num)
	fmt.Println(a)
}
func solution(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		least := i
		for j := i; j < len(array); j++ {
			if array[j] < array[least] {
				least = j
			}
		}
		if i != least {
			array[i], array[j] = array[j], array[i]
		}
	}
	return array
}
