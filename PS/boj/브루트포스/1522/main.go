package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	str string
)

func solve() {
	
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &str)


}