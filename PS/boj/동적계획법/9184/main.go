package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	dp     [21][21][21]int
)

func w(a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}

	if a > 20 || b > 20 || c > 20 {
		return w(20, 20, 20)
	}

	ret := &dp[a][b][c]
	if *ret != -1 {
		return *ret
	}

	if a < b && b < c {
		*ret = w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
	} else {
		*ret = w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
	}

	return *ret
}

func main() {
	defer writer.Flush()

	for i := 0; i < 21; i++ {
		for j := 0; j < 21; j++ {
			for k := 0; k < 21; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	for {
		line, _, _ := reader.ReadLine()
		nums := strings.Fields(string(line))
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		c, _ := strconv.Atoi(nums[2])

		if a == -1 && b == -1 && c == -1 {
			break
		}

		ans := w(a, b, c)

		fmt.Fprintf(writer, "w(%d, %d, %d) = %d\n", a, b, c, ans)
	}
}
