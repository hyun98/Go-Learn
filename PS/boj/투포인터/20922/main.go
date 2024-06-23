package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	N      int
	K      int
	arr    []int
	count  [100001]int
	answer = 1
)

func solve() {
	lo, hi := 0, 1
	count[arr[0]]++
	length := 1

	for lo <= hi && hi < N {
		if count[arr[hi]]+1 > K {
			count[arr[lo]]--
			lo++
			length--
		} else {
			count[arr[hi]]++
			hi++
			length++
			answer = max(answer, length)
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N, &K)

	for i := 0; i < N; i++ {
		a := 0
		fmt.Fscan(reader, &a)
		arr = append(arr, a)
	}

	solve()

	fmt.Fprint(writer, answer)

}
