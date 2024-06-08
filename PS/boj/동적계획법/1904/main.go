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
	MOD    = 15746
)

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	num, _ := strconv.Atoi(string(line))

	a, b := 1, 2

	for i := 1; i < num; i++ {
		a, b = b, (a+b)%MOD
	}
	
	fmt.Print(a)
}
