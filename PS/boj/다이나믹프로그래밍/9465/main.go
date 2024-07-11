package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	T, N    int
	cache   = [2][100001]int{}
	sticker = [2][100001]int{}
)

func solve(r, c int) int {
	if c >= N {
		return 0
	}

	ret := &cache[r][c]
	if *ret != -1 {
		return *ret
	}

	*ret = sticker[r][c] + int(math.Max(float64(solve(1-r, c+1)), float64(solve(1-r, c+2))))
	return *ret
}

func resetCache() {
	for i := 0; i < 2; i++ {
		for j := 0; j < N; j++ {
			cache[i][j] = -1
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &T)
	for ; T > 0; T-- {
		fmt.Fscan(reader, &N)
		resetCache()
		for i := 0; i < 2; i++ {
			for j := 0; j < N; j++ {
				fmt.Fscan(reader, &sticker[i][j])
			}
		}

		ansA := solve(0, 0)
		ansB := solve(1, 0)

		ans := int(math.Max(float64(ansA), float64(ansB)))

		fmt.Fprintln(writer, ans)
	}
}
