package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	N      int
	ans    int
	cqueen [16]int
)

func canAttack(r int) bool {
	for i := 0; i < r; i++ {
		if cqueen[i] == cqueen[r] {
			return true
		}
		if math.Abs(float64(i-r)) == math.Abs(float64(cqueen[i]-cqueen[r])) {
			return true
		}
	}

	return false
}

func nQueen(r int) {
	if r == N {
		ans++
		return
	}

	for i := 0; i < N; i++ {
		cqueen[r] = i

		if canAttack(r) {
			continue
		}

		nQueen(r + 1)
	}
}

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	N, _ = strconv.Atoi(string(line))

	nQueen(0)

	writer.WriteString(strconv.Itoa(ans))
}
