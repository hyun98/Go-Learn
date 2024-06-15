package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	str     string
	visited [51]bool
	found   bool
)

func StrToInt(st, ed int) int {
	ret, _ := strconv.Atoi(str[st:ed])
	return ret
}

func Solve(n int, ans *[]int) {
	if found {
		return
	}

	if n == len(str) {
		mv := slices.Max(*ans)

		if len(*ans) == mv {
			found = true
			for _, a := range *ans {
				fmt.Fprint(writer, a, " ")
			}
		}

		return
	}

	// 1자리
	num := StrToInt(n, n+1)

	if num == 0 {
		return
	}

	if !visited[num] {
		visited[num] = true
		*ans = append(*ans, num)
		Solve(n+1, ans)
		*ans = (*ans)[:len(*ans)-1]
		visited[num] = false
	}

	// 2자리
	if n+2 > len(str) {
		return
	}

	num = StrToInt(n, n+2)

	if num > 50 {
		return
	}

	if !visited[num] {
		visited[num] = true
		*ans = append(*ans, num)
		Solve(n+2, ans)
		*ans = (*ans)[:len(*ans)-1]
		visited[num] = false
	}
}

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &str)

	Solve(0, &[]int{})
}
