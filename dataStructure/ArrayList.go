package main

import "fmt"

func main() {
	var ArrayList = []int{1, 2, 3, 4}

	ArrayList = append(ArrayList, 5)

	fmt.Println(ArrayList, len(ArrayList))
}
