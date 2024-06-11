package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader    = bufio.NewReader(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N, M, K   int
	board     [2001][2001]byte
	aType     = []string{"BW", "WB"}
	bType     = []string{"WB", "BW"}
	aSumBoard [2001][2001]int
	bSumBoard [2001][2001]int
	answer    = float64(2000 * 2000)
)

func makeSumBoard() {
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			var va int = 0
			var vb int = 0
			if board[i-1][j-1] != aType[(i-1)%2][(j-1)%2] {
				va = 1
			}

			if board[i-1][j-1] != bType[(i-1)%2][(j-1)%2] {
				vb = 1
			}

			aSumBoard[i][j] = va + aSumBoard[i-1][j] + aSumBoard[i][j-1] - aSumBoard[i-1][j-1]
			bSumBoard[i][j] = vb + bSumBoard[i-1][j] + bSumBoard[i][j-1] - bSumBoard[i-1][j-1]
		}
	}
}

func calculate() {
	for i := 0; i < N-K+1; i++ {
		for j := 0; j < M-K+1; j++ {
			amin := aSumBoard[i+K][j+K] - aSumBoard[i][j+K] - aSumBoard[i+K][j] + aSumBoard[i][j]
			bmin := bSumBoard[i+K][j+K] - bSumBoard[i][j+K] - bSumBoard[i+K][j] + bSumBoard[i][j]
			answer = math.Min(answer, float64(amin))
			answer = math.Min(answer, float64(bmin))
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N, &M, &K)

	var line string

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &line)
		for j := 0; j < M; j++ {
			board[i][j] = line[j]
		}
	}

	makeSumBoard()
	calculate()

	fmt.Fprint(writer, int(answer))
}
