package main

import (
	"fmt"
	"time"
)

func main() {
	go func1()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second / 10)
		fmt.Println(i)
	}
}

func func1() {
	for i := 0; i <= 20; i++ {
		time.Sleep(time.Second / 10)
		fmt.Println(i)
	}
}
