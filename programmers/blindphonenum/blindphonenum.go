package main

import (
	"fmt"
)

func main() {
	tr := solution("01039083375")
	fmt.Println(tr)
}

func solution(phone_number string) string {
	var extra string
	for i := 0; i < len(phone_number)-4; i++ {
		extra = extra + "*"
	}
	phone := phone_number[len(phone_number)-4:] //3375 젤로 마지막부터

	return extra + phone
}
