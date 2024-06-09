package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	psum [200002][27]int
	str  string
	N    int
)

func makePsum() {
	for i := 1; i <= len(str); i++ {
		for q := 0; q < 26; q++ {
			if q == int(str[i-1]-'a') {
				psum[i][q] = psum[i-1][q] + 1
			} else {
				psum[i][q] = psum[i-1][q]
			}
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &str, &N)

	makePsum()

	ans := []int{}

	for i := 0; i < N; i++ {
		var target string
		var sIdx int
		var eIdx int
		fmt.Fscan(reader, &target, &sIdx, &eIdx)
		ans = append(ans, psum[eIdx+1][target[0]-'a']-psum[sIdx][target[0]-'a'])
	}

	for i := range ans {
		fmt.Fprintf(writer, "%d\n", ans[i])
	}
}
