package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.WriteString("\\    /\\\n")
	writer.WriteString(" )  ( ')\n")
	writer.WriteString("(  /  )\n")
	writer.WriteString(" \\(__)|\n")
}
