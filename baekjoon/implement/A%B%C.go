package main

import "fmt"

func main() {
	A, B, C := 0, 0, 0
	fmt.Scanf("%d %d %d", &A, &B, &C)
	fmt.Print((A+B)%C, "\n", ((A%C)+(B%C))%C, "\n", (A*B)%C, "\n", ((A%C)*(B%C))%C)
}
