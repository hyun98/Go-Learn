package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	line, _, _ := reader.ReadLine()
	nums := strings.Fields(string(line))
	h, _ := strconv.Atoi(nums[0])
	m, _ := strconv.Atoi(nums[1])

	line2, _, _ := reader.ReadLine()
	t, _ := strconv.Atoi(string(line2))

	daymin := 1440
	hourmin := 60

	startTime := h*hourmin + m

	endTime := (startTime + t) % daymin
	endHour := endTime / hourmin
	endMin := endTime % hourmin

	output := fmt.Sprintf("%d %d\n", endHour, endMin)
	writer.WriteString(output)
}
