package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	M       int
	visited [9]bool
)

func backTrack(depth int, ans *[]string) {
	if depth == M {
		writer.WriteString(strings.Join(*ans, " ") + "\n")
		return
	}

	for i := 1; i <= N; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		*ans = append(*ans, strconv.Itoa(i))
		backTrack(depth+1, ans)
		*ans = (*ans)[:len(*ans)-1]
		visited[i] = false
	}
}

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	nums := strings.Fields(string(line))
	N, _ = strconv.Atoi(nums[0])
	M, _ = strconv.Atoi(nums[1])

	ans := []string{}
	backTrack(0, &ans)
}
