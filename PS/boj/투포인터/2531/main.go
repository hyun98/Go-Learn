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
	d      int
	k      int
	c      int
	sushi  []int
	count  [3001]int
	answer int
)

func solve() {
	num := 0
	for i := 0; i < k; i++ {
		if count[sushi[i]] == 0 {
			num++
		}
		count[sushi[i]]++
	}

	if count[c] == 0 {
		num++
		count[c]++
	}

	answer = max(answer, num)

	for i := k; i < len(sushi); i++ {
		count[sushi[i-k]]--
		if count[sushi[i-k]] == 0 {
			num--
		}

		if count[sushi[i]] == 0 {
			num++
		}
		count[sushi[i]]++

		if count[c] == 0 {
			num++
			count[c]++
		}
		answer = max(answer, num)
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N, &d, &k, &c)
	for i := 0; i < N; i++ {
		a := 0
		fmt.Fscan(reader, &a)
		sushi = append(sushi, a)
	}

	for i := 0; i < k-1; i++ {
		sushi = append(sushi, sushi[i])
	}

	solve()

	fmt.Fprint(writer, answer)
}
