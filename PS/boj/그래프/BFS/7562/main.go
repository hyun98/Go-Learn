package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Loc struct {
	r, c int
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	T      int
	N      int
	start  Loc
	end    Loc
	dr     = [8]int{-2, -1, 1, 2, 2, 1, -1, -2}
	dc     = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
	INF    = 1<<31 - 1
)

func solve() int {
	dp := [301][301]int{}

	// clear
	for i := 0; i < 301; i++ {
		for j := 0; j < 301; j++ {
			dp[i][j] = INF
		}
	}

	que := []Loc{start}
	dp[start.r][start.c] = 0

	for len(que) > 0 {
		now := que[0]
		que = que[1:]

		for i := 0; i < 8; i++ {
			nr := now.r + dr[i]
			nc := now.c + dc[i]

			if nr >= N || nr < 0 || nc >= N || nc < 0 ||
				dp[now.r][now.c]+1 >= dp[nr][nc] {
				continue
			}

			que = append(que, Loc{nr, nc})
			dp[nr][nc] = dp[now.r][now.c] + 1
		}
	}

	return dp[end.r][end.c]
}

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	T, _ = strconv.Atoi(string(line))

	for i := 0; i < T; i++ {
		line, _, _ = reader.ReadLine()
		N, _ = strconv.Atoi(string(line))

		line, _, _ = reader.ReadLine()
		nums := strings.Fields(string(line))
		start.r, _ = strconv.Atoi(nums[0])
		start.c, _ = strconv.Atoi(nums[1])

		line, _, _ = reader.ReadLine()
		nums = strings.Fields(string(line))
		end.r, _ = strconv.Atoi(nums[0])
		end.c, _ = strconv.Atoi(nums[1])

		answer := solve()
		writer.WriteString(strconv.Itoa(answer) + "\n")
	}
}
