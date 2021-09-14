/*
보험 항목별 보장 금액이 순서대로 들어있는 배열이 있습니다. 주어진 항목 중에서 k개 이상의 항목을 선택하여 새 보험 상품을 만들되, 새 보험 상품에 포함되는 각 항목의 보장 금액의 합이 t 이하가 되도록 하려고 합니다. 다음은 항목별 보장 금액이 들어있는 배열 [2, 5, 3, 8, 1], k = 3, t = 11 이 주어진 경우의 예시입니다.
먼저 주어진 배열은 5개의 보험 항목의 보장 금액이 순서대로 [2, 5, 3, 8, 1]임을 나타냅니다. 이때, 3개 이상의 항목을 선택하는 방법은 총 16가지가 있는데, 이 중, 선택된 항목들의 보장 금액의 합이 11 이하가 되는 경우는 다음과 같이 6가지가 있습니다.

[2, 5, 3]
[2, 5, 1]
[5, 3, 1]
[2, 3, 1]
[2, 8, 1]
[2, 5, 3, 1]
예를 들어 선택한 항목의 보장 금액이 [2, 5, 3, 1] 인 경우 보장 금액의 합은 2 + 5 + 3 + 1 = 11이므로 11 이하가 됩니다. 그러나 만약 선택한 항목의 보장 금액이 [5, 3, 8]인 경우 보장 금액의 합은 5 + 3 + 8 = 16이므로 이 경우는 불가능한 방법이 됩니다.
항목별 보장 금액이 순서대로 들어있는 배열 arr와 k, t가 매개변수로 주어질 때, arr에서 k개 이상의 항목을 선택했을 때, 금액 합이 t 이하가 되도록 하는 경우의 수를 return 하도록 solution 함수를 완성해주세요.

제한사항
arr는 각 보험 항목별 보장 금액이 순서대로 들어있는 배열이며, 길이는 1 이상 15 이하 입니다.
arr의 각 원소는 1 이상 100,000 이하의 자연수입니다.
k는 1 이상 15 이하의 자연수이며, 항상 arr 의 길이 이하인 경우만 주어집니다.
t는 1 이상 1,000,000 이하의 자연수입니다.
입출력 예
arr	k	t	result
[2,5,3,8,1]	3	11	6
[1,1,2,2]	2	3	5
[1,2,3,2]	2	2	0
입출력 예 설명
입출력 예 #1
문제의 예시와 같습니다.

입출력 예 #2
다음과 같이 5가지 경우가 가능합니다.

[1, 1], (첫 번째, 두 번째 항목)
[1, 2], (첫 번째, 세 번째 항목)
[1, 2], (첫 번째, 네 번째 항목)
[1, 2], (두 번째, 세 번째 항목)
[1, 2], (두 번째, 네 번째 항목)
입출력 예 #3
2가지 이상을 선택하면서 금액 합계가 2 이하가 되도록 하는 방법은 없습니다.
*/
package main

import (
	"container/list"
	"fmt"
)

type Combination struct {
	n    int
	c    int
	list *list.List
}

func (comb *Combination) setParams(N, C int) *Combination {
	(*comb).n = N
	(*comb).c = C
	(*comb).list = list.New()
	return comb
}

func (comb *Combination) getLists() *list.List {
	return (*comb).list
}

func (comb *Combination) makeCombination() {
	selected := make([]bool, (*comb).n)
	comb.dfs(0, 0, (*comb).c, selected)
}

func (comb *Combination) dfs(idx, cnt, target int, selected []bool) {
	if cnt == target {
		result := make([]bool, (*comb).n)
		for i := 0; i < len(result); i++ {
			result[i] = selected[i]
		}
		comb.list.PushBack(result)
		return
	}
	for i := idx; i < (*comb).n; i++ {
		if selected[i] {
			continue
		}
		selected[i] = true
		comb.dfs(i, cnt+1, target, selected)
		selected[i] = false
	}
}

func New() Combination {
	return Combination{
		n:    0,
		c:    0,
		list: list.New(),
	}
}
func solution(arr []int, k int, t int) int {
	c := New()

	c.setParams(5, 3)

	c.makeCombination()

	first := c.list.Front()
	for first.Next() != nil {
		fmt.Println(first.Value)
		first = first.Next()
	}

	return -1
}

func main() {
	arr := []int{2, 5, 3, 8, 1}
	k := 3
	t := 11
	a := solution(arr, k, t)
	fmt.Println(a)
}
