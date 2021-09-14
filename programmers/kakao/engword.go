package main

import (
	"fmt"
	"strings"
)

func solution(s string) int {

	var a, b int
	a = strings.Index(s, "zero")
	b = strings.LastIndex(s, "one")

	fmt.Println(a, b)
	return a

}
func main() {
	s := "one4seveneightzero"
	a := solution(s)
	fmt.Println(a)
}
