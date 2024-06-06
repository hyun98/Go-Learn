package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	mp = make(map[string]bool)
)

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	str := string(line)

	for i := 1; i <= len(str); i++ {
		for s := 0; s + i <= len(str); s++ {
			part := str[s : s + i]
			mp[part] = true
		}
	}

	writer.WriteString(strconv.Itoa(len(mp)))
}