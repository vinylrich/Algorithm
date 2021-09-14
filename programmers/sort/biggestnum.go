package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{6, 10, 2}
	a := solution(numbers)
	fmt.Println(a)
}

func solution(numbers []int) string {
	var strarray []string
	var tmpstr string
	var rst string
	for i := 0; i < len(numbers); i++ {
		tmpstr = strconv.Itoa(numbers[i])
		strarray = append(strarray, tmpstr)
	} //{"6", "10","2"}
	for i := 0; i < len(numbers); i++ { //3번 돌아감
		for j := 0; j < len(numbers)-i-1; j++ {
			a, _ := strconv.Atoi(strarray[j] + strarray[j+1]) //330
			b, _ := strconv.Atoi(strarray[j+1] + strarray[j]) //303
			if a < b {
				strarray[j], strarray[j+1] = strarray[j+1], strarray[j] //6 10
			} //[3, 30, 34, 5, 9]
		}
	}
	rst = strings.Join(strarray, "")
	return rst
}
