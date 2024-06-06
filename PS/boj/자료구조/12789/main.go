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
	N      int
	lines  []int //queue
	stack  []int //stack
)

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	N, _ = strconv.Atoi(string(line))

	line, _, _ = reader.ReadLine()
	nums := strings.Fields(string(line))

	for _, n := range nums {
		a, _ := strconv.Atoi(n)
		lines = append(lines, a)
	}

	now := 1
	for len(lines) != 0 || len(stack) != 0 {
		if len(lines) > 0 && len(stack) == 0 {
			d := lines[0]
			lines = lines[1:]
			if d == now {
				now++
			} else {
				stack = append(stack, d)
			}
		} else if len(lines) > 0 && len(stack) > 0 {
			a := lines[0]
			b := stack[len(stack)-1]

			if a == now {
				lines = lines[1:]
				now++
			} else if b == now {
				stack = stack[:len(stack)-1]
				now++
			} else {
				lines = lines[1:]
				stack = append(stack, a)
			}

		} else if len(lines) == 0 && len(stack) > 0 {
			b := stack[len(stack)-1]
			if b == now {
				stack = stack[:len(stack)-1]
				now++
			} else {
				writer.WriteString("Sad")
				return
			}
		}
	}

	writer.WriteString("Nice")
}
