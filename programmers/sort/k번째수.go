package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	a = []int{1, 5, 2, 6, 3, 7, 4}
	var commands [][]int
	commands = [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}
	rst := solution(a, commands)
	fmt.Println(rst)

}
func solution(array []int, commands [][]int) []int {
	var rst []int

	for i := 0; i < len(commands); i++ {
		tmp1 := make([]int, len(array))
		copy(tmp1, array)
		fmt.Println(tmp1, array)

		var cnt int = commands[i][2] - 1               //몇번째 값을 추출할지 2번인덱스를 추출한다
		tmp1 = tmp1[commands[i][0]-1 : commands[i][1]] //1번부터 4번까지 3번부터 4번까지

		sort.Slice(tmp1, func(i2, j int) bool {
			return tmp1[i2] < tmp1[j]
		})

		fmt.Println(tmp1)
		rst = append(rst, tmp1[cnt]) //arr 2번인덱스를 넣기
	} //1:4
	return rst
}
