package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x int
	fmt.Scanln(&x)
	a := solution(x)
	fmt.Println(a)
}

func solution(x int) bool {
	var sum, tmp int
	str := strconv.Itoa(x)
	strar := strings.Split(str, "")
	for _, p := range strar {
		tmp, _ = strconv.Atoi(p)
		sum += tmp
	}

	if x%sum == 0 {
		return true
	}
	return false
}
