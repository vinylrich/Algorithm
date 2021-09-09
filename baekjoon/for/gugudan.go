package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}
}
