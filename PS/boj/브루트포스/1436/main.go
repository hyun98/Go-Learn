package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	N      int
	title  []string
)

func main() {
	defer writer.Flush()

	line, _, _ := reader.ReadLine()
	nums := strings.Fields(string(line))
	N, _ = strconv.Atoi(nums[0])

	start := 666

	for i := 0; len(title) < N; i++ {
		st := strconv.Itoa(start)
		if strings.Contains(st, "666") {
			title = append(title, st)
		}
		start++
	}

	writer.WriteString(title[N-1])
}
