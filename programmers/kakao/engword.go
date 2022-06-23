package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(s string) int {

	//f := strings.NewReplacer(
	//	"zero", "0",
	//	"one", "1",
	//	"two", "2",
	//	"three", "3",
	//	"four", "4",
	//	"five", "5",
	//	"six", "6",
	//	"seven", "7",
	//	"eight", "8",
	//	"nine", "9",
	//)
	//i, _ := strconv.Atoi(f.Replace(s))
	//return i
	s = strings.Replace(s, "zero", "0", -1)
	s = strings.Replace(s, "one", "1", -1)
	s = strings.Replace(s, "two", "2", -1)
	s = strings.Replace(s, "three", "3", -1)
	s = strings.Replace(s, "four", "4", -1)
	s = strings.Replace(s, "five", "5", -1)
	s = strings.Replace(s, "six", "6", -1)
	s = strings.Replace(s, "seven", "7", -1)
	s = strings.Replace(s, "eight", "8", -1)
	s = strings.Replace(s, "nine", "9", -1)
	convInt, _ := strconv.Atoi(s)
	return convInt
}
func main() {
	s := "one4seveneightzero"
	fmt.Println(solution(s))
}
