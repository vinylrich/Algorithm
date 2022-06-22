package main

import "fmt"

func main() {
	fmt.Println(solution(5, 24))
}
func solution(a int, b int) string {
	var cntmonth int
	var monthSum int
	for i := 1; i <= a; i++ {
		switch i {
		case 1:
			cntmonth = 0
		case 2:
			cntmonth = 31
		case 3:
			cntmonth = 29
		case 4:
			cntmonth = 31
		case 5:
			cntmonth = 30
		case 6:
			cntmonth = 31
		case 7:
			cntmonth = 30
		case 8:
			cntmonth = 31
		case 9:
			cntmonth = 31
		case 10:
			cntmonth = 30
		case 11:
			cntmonth = 31
		case 12:
			cntmonth = 30
		}
		monthSum += cntmonth
	}
	weekRst := (monthSum+b)%7 - 1
	switch weekRst {
	case 0:
		return "FRI"
	case 1:
		return "SAT"
	case 2:
		return "SUN"
	case 3:
		return "MON"
	case 4:
		return "TUE"
	case 5:
		return "WED"
	case 6:
		return "THU"
	case -1:
		return "THU"
	}
	return "Something is Wrong"
}

//func solution(a int, b int) string {
//	num2day := map[int]string {
//		0: "SUN",
//		1: "MON",
//		2: "TUE",
//		3: "WED",
//		4: "THU",
//		5: "FRI",
//		6: "SAT",
//	}
//	date := 5
//	for month := 1; month < a; month++ {
//		switch month {
//		case 1, 3, 5, 7, 8, 10, 12: date += 31
//		case 4, 6, 9, 11: date += 30
//		default: date += 29
//		}
//	}
//	return num2day[(date + b - 1) % 7]
//}
