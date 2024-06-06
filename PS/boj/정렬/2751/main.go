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
	N      int
	arr    []int
)

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	N, _ = strconv.Atoi(string(line))

	for i := 0; i < N; i++ {
		line, _, _ := reader.ReadLine()
		a, _ := strconv.Atoi(string(line))
		arr = append(arr, a)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for _, val := range arr {
		writer.WriteString(strconv.Itoa(val) + "\n")
	}
}
