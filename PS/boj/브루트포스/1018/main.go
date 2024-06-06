package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	mp     [51]string
	N      int
	M      int
	INF    = int((^uint(0)) >> 1)
	A      = [8]string{"W", "B", "W", "B", "W", "B", "W", "B"}
	B      = [8]string{"B", "W", "B", "W", "B", "W", "B", "W"}
)

func Check(r, c int) int {
	retA := 0
	retB := 0

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i%2 == 0 {
				if string(mp[i+r][j+c]) != A[j] {
					retA++
				}
			} else {
				if string(mp[i+r][j+c]) != B[j] {
					retA++
				}
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i%2 == 0 {
				if string(mp[i+r][j+c]) != B[j] {
					retB++
				}
			} else {
				if string(mp[i+r][j+c]) != A[j] {
					retB++
				}
			}
		}
	}

	return min(retA, retB)
}

func main() {
	defer writer.Flush()

	line, _, _ := reader.ReadLine()
	nums := strings.Fields(string(line))
	N, _ = strconv.Atoi(nums[0])
	M, _ = strconv.Atoi(nums[1])

	for i := 0; i < N; i++ {
		line, _, _ := reader.ReadLine()
		mp[i] = string(line)
	}

	answer := 64

	for r := 0; N-r >= 8; r++ {
		for c := 0; M-c >= 8; c++ {
			answer = min(answer, Check(r, c))
		}
	}

	writer.WriteString(strconv.Itoa(answer))
}
