package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("%")
		}
		fmt.Println(" ")
	}
}
