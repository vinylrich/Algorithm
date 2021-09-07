package main

import (
	"fmt"
)

func main() {
	a, b := 0.0, 0.0
	fmt.Scanf("%f %f", &a, &b)

	fmt.Printf("%.9f", a/b)
}
