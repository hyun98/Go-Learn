package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	a, _, _ := reader.ReadLine()
	n, _ := strconv.Atoi(string(a))

	for i := 0; i < n; i++ {
		a, _, _ := reader.ReadLine()
		fields := strings.Fields(string(a))
		b, _ := strconv.Atoi(fields[0])
		c, _ := strconv.Atoi(fields[1])
		writer.WriteString(strconv.Itoa(b+c) + "\n")
	}
}
