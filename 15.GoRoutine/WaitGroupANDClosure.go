package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU Num: ", runtime.NumCPU())
	wait := new(sync.WaitGroup)
	wait.Add(102)

	str := "BlockChain"

	go func() {
		defer wait.Done()
		fmt.Println("hello")
	}()

	go func() {
		defer wait.Done()
		fmt.Println(str)
	}()

	for i := 0; i < 100; i++ {
		go func(n int) {
			defer wait.Done()
			fmt.Println(n)
		}(i)
	}

	wait.Wait()
}
