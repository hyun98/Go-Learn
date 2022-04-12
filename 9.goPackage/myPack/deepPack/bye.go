package deepPack

import "fmt"

func Goodbye() func() {
	count := 0
	return func() {
		count++
		fmt.Println(count, "GoodBye!")
	}
}

func noGoodbye() {
	fmt.Println("noGoodbye")
}
