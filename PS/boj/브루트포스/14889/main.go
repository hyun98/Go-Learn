package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	board   [21][21]int
	visited [21]bool
	ans     = math.MaxInt
)

func getStat(team *[]int) (int, int) {
	statA, statB := 0, 0
	teamB := []int{}
	for i := 0; i < N; i++ {
		if visited[i] {
			continue
		}
		teamB = append(teamB, i)
	}

	for i := 0; i < len(*team)-1; i++ {
		for j := i + 1; j < len(*team); j++ {
			statA += (board[(*team)[i]][(*team)[j]] + board[(*team)[j]][(*team)[i]])
			statB += (board[teamB[i]][teamB[j]] + board[teamB[j]][teamB[i]])
		}
	}

	return statA, statB
}

func makeTeam(now int, teamA *[]int) {
	if len(*teamA) == N/2 {
		// 능력치 구하기
		statA, statB := getStat(teamA)
		diff := math.Abs(float64(statA) - float64(statB))

		// 최솟값 비교하기
		ans = int(math.Min(float64(ans), diff))
		return
	}

	for i := now; i < N; i++ {
		visited[i] = true
		*teamA = append(*teamA, i)
		makeTeam(i+1, teamA)
		*teamA = (*teamA)[:len(*teamA)-1]
		visited[i] = false
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			var r int
			fmt.Fscan(reader, &r)
			board[i][j] = r
		}
	}

	makeTeam(0, &[]int{})

	fmt.Fprint(writer, ans)
}
