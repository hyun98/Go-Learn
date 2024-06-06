package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	line, _, _ := reader.ReadLine()

	var arr []int
	for _, c := range string(line) {
		arr = append(arr, int(c-'0'))
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	ans := ""

	for _, i := range arr {
		ans += strconv.Itoa(i)
	}

	writer.WriteString(ans)
}
