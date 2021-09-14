package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N int
	var A []int
	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)

	N, _ = strconv.Atoi(str)
	A = make([]int, N)
	str, _ = r.ReadString('\n')
	strTkn := make([]string, N)
	for i := 0; i < N; i++ {
		strTkn = strings.Split(str, " ")

		if i == N-1 {
			strTkn[i] = strings.TrimSpace(strTkn[i])
		}
		A[i], _ = strconv.Atoi(strTkn[i])
	}
	sort.Sort(sort.IntSlice(A))

	max, min := A[0], A[N-1]

	w.WriteString(strconv.Itoa(max) + " " + strconv.Itoa(min))

	w.Flush()
}
