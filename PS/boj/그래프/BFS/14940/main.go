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
	M      int
	Board  [1001][1001]int
	dist   [1001][1001]int
	dr     = [4]int{0, 0, -1, 1}
	dc     = [4]int{1, -1, 0, 0}
	Start  Loc
)

type Loc struct {
	r, c, d int
}

func bfs() {
	que := []Loc{}
	que = append(que, Start)
	dist[Start.r][Start.c] = Start.d

	for len(que) > 0 {
		now := que[0]
		que = que[1:]
		ndist := dist[now.r][now.c]
		// 같은 경우가 가장 최단거리
		if ndist != now.d {
			continue
		}

		for i := 0; i < 4; i++ {
			nr := now.r + dr[i]
			nc := now.c + dc[i]
			if nr < 0 || nc < 0 || nr >= N || nc >= M || Board[nr][nc] == 0 {
				continue
			}
			if ndist+1 < dist[nr][nc] {
				dist[nr][nc] = ndist + 1
				que = append(que, Loc{nr, nc, ndist + 1})
			}
		}
	}
}

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &N, &M)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fscan(reader, &Board[i][j])
			if Board[i][j] == 2 {
				Start = Loc{i, j, 0}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			dist[i][j] = 1<<31 - 1
		}
	}

	bfs()

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if dist[i][j] == 1<<31-1 {
				if Board[i][j] == 0 {
					dist[i][j] = 0
				} else {
					dist[i][j] = -1
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fprint(writer, dist[i][j], " ")
		}
		fmt.Fprintln(writer)
	}
}
