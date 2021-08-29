package main

import "fmt"

func main() {
	priorities := []int{2, 1, 3, 2}
	location := 2

	var a *int
	a = &priorities[location] //주소 0x0001

	fmt.Println(a, &priorities[location])

}
