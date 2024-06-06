package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	bytes, _, _ := reader.ReadLine()
	line := strings.ToUpper(string(bytes))

	mp := &map[rune]int{}

	for _, r := range line {
		(*mp)[r]++
	}

	maxCnt := 0
	var fqchar []rune
	for char, count := range *mp {
		if count > maxCnt {
			maxCnt = count
			fqchar = []rune{char}
		} else if count == maxCnt {
			fqchar = append(fqchar, char)
		}
	}

	if len(fqchar) > 1 {
		fmt.Fprintln(writer, "?")
	} else {
		fmt.Fprintf(writer, "%c\n", fqchar[0])
	}
}
