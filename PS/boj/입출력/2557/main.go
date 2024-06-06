package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	str := 'b'
	fmt.Println(str + 'a')

	fmt.Println(int32((^uint32(0)) >> 1))
}
