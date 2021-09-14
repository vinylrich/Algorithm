package main

func main() {

}
func solution(arr []int) float64 {
	var sum float64
	for _, p := range arr {
		sum += float64(p)
	}
	return sum
}
