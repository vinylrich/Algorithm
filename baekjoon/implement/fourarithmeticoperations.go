package main

import "fmt"

func main() {
	a, b := 0, 0

	fmt.Scanf("%d %d", &a, &b)
	fmt.Print(a+b, "\n", a-b, "\n", a*b, "\n", a/b, "\n", a%b, "\n")
}
