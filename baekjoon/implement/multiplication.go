package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d\n%d", &a, &b)

	third := a * (b % 10)
	fourth := a * ((b / 10) % 10)
	fifth := a * (b / 100)
	rst := a * b
	fmt.Print(third, "\n", fourth, "\n", fifth, "\n", rst)
}
