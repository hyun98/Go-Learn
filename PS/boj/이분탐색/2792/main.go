package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	N      int
	M      int
	C      []int
)

func bi(mid int) int {
	ret := 0
	for _, c := range C {
		ret += (c / mid)
		if m := c % mid; m != 0 {
			ret++
		}
	}
	return ret
}

func solve() int {
	lo := 1
	hi := slices.Max(C)
	mid := 0
	ret := 0

	for lo <= hi {
		mid = (lo + hi) / 2
		temp := bi(mid)
		if temp <= N {
			hi = mid - 1
			ret = mid
		} else {
			lo = mid + 1
		}
	}

	return ret
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N, &M)

	for i := 0; i < M; i++ {
		a := 0
		fmt.Fscan(reader, &a)
		C = append(C, a)
	}

	ans := solve()

	fmt.Fprint(writer, ans)
}
