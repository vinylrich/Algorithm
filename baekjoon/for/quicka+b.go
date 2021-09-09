package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)

	N, _ := strconv.Atoi(str)

	for i := 0; i < N; i++ {
		str, _ = r.ReadString('\n')
		str = strings.TrimSpace(str)
		strTkn := strings.Split(str, " ")

		a, _ := strconv.Atoi(strTkn[0])
		b, _ := strconv.Atoi(strTkn[1])

		res := a + b
		w.WriteString(strconv.Itoa(res) + "\n")
	}

	w.Flush()
}
