package main

import "fmt"

func main() {
	grade := 0

	fmt.Scanf("%d", &grade)
	if grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {

		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else if grade >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
