package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)

	n, _ := strconv.Atoi(str)

	for i := 1; i <= n; i++ {
		sum += i
	}
	w.WriteString(strconv.Itoa(sum))
	w.Flush()
}
