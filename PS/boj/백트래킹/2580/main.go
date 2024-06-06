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
	mp     [9][9]int
	fin    = false
)

func check(r, c, num int) bool {
	// r행 확인
	for i := 0; i < 9; i++ {
		if mp[r][i] == num {
			return false
		}
	}

	// c열 확인
	for i := 0; i < 9; i++ {
		if mp[i][c] == num {
			return false
		}
	}

	// 3x3 사각형 확인
	for i := (r / 3) * 3; i < (r/3)*3+3; i++ {
		for j := (c / 3) * 3; j < (c/3)*3+3; j++ {
			if mp[i][j] == num {
				return false
			}
		}
	}

	return true
}

func back(r, c int) {
	if fin {
		return
	}

	if r == 9 {
		fin = true
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Print(strconv.Itoa(mp[i][j]) + " ")
			}
			fmt.Println()
		}
		return
	}

	if c == 9 {
		back(r+1, 0)
		return
	}

	if mp[r][c] != 0 {
		back(r, c+1)
	} else {
		for i := 1; i <= 9; i++ {
			if check(r, c, i) {
				mp[r][c] = i
				back(r, c+1)
			}
			mp[r][c] = 0
		}
	}
}

func main() {
	defer writer.Flush()
	for r := 0; r < 9; r++ {
		line, _, _ := reader.ReadLine()
		row := strings.Fields(string(line))
		for c, s := range row {
			mp[r][c], _ = strconv.Atoi(string(s))
		}
	}

	back(0, 0)
}
