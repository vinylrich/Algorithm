package main

import "fmt"

func main() {
	var hour, minute, extra int
	fmt.Scan(&hour, &minute)
	if minute >= 45 {
		minute = minute - 45
	} else {
		extra = 45 - minute
		hour--
		if hour < 0 {
			hour = 23
		}
		minute = 60 - extra
	}
	fmt.Println(hour, minute)
}
