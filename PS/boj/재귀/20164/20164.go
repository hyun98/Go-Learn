package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	N      string
	minAns = (1 << 31) - 1
	maxAns = 0
)

func getOddCount(num string) int {
	ret := 0
	for _, n := range num {
		if int(n-'0')%2 == 1 {
			ret++
		}
	}
	return ret
}

func devide(num string, ans int) {
	if len(num) == 1 {
		minAns = min(minAns, ans)
		maxAns = max(maxAns, ans)
		return
	}

	if len(num) == 2 {
		a, _ := strconv.Atoi(num[:1])
		b, _ := strconv.Atoi(num[1:])
		nn := strconv.Itoa(a + b)
		devide(nn, ans+getOddCount(nn))
		return
	}

	for i := 1; i < len(num); i++ {
		for j := 2; j < len(num); j++ {
			if i >= j {
				continue
			}
			a, _ := strconv.Atoi(num[:i])
			b, _ := strconv.Atoi(num[i:j])
			c, _ := strconv.Atoi(num[j:])
			nn := strconv.Itoa(a + b + c)
			devide(nn, ans+getOddCount(nn))
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)

	devide(N, getOddCount(N))

	fmt.Fprint(writer, minAns, maxAns)
}
