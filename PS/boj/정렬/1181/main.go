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
	words  []string
)

func makeSliceUnique(s []string) []string {
	mp := make(map[string]struct{})
	ret := []string{}

	for _, val := range s {
		if _, ok := mp[val]; !ok {
			mp[val] = struct{}{}
			ret = append(ret, val)
		}
	}
	return ret
}

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	N, _ = strconv.Atoi(string(line))

	for i := 0; i < N; i++ {
		line, _, _ := reader.ReadLine()
		words = append(words, string(line))
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		} else if len(words[i]) > len(words[j]) {
			return false
		} else {
			return words[i] < words[j]
		}
	})

	words = makeSliceUnique(words)

	for _, w := range words {
		writer.WriteString(w + "\n")
	}
}
