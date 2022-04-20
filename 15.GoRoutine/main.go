package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// go 키워드를 사용하여 함수를 호출하면 런타임에서 새로운 goroutine 실행한다.
// Go 루틴은 OS 스레드보다 훨씬 가볍게 비동기 처리를 구현하기 위해 만들어졌다.
// Go 루틴들은 OS 스레드와 1 대 1 로 대응되지 않고, Multiplexing 으로 훨씬 적은 OS 를 사용한다.

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "****", i)
	}
}

func main() {
	// 비동기/동기 차이
	// say 가 0 부터 9까지 출력하고
	// Async say 3개가 불규칙적으로 출력한다.
	say("func")
	go say("Async1")
	go say("Async2")
	go say("Async3")
	time.Sleep(time.Second * 1)

	// WaitGroup 생성. 2개의 Go 루틴을 기다린다.
	var wait sync.WaitGroup
	wait.Add(2)

	// 람다 사용 goroutine
	go func() {
		defer wait.Done() // 끝나면 Done() 호출
		fmt.Println("Hello")
	}()

	// 람다함수에 파라미터 전달
	str := "Hi"
	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}(str)

	wait.Wait() // Go 루틴 모두 끝날 때까지 대기

	// 기기가 복수개의 CPU 를 가지는 경우에, 다중 CPU 에서 병렬처리를 지정해줄 수 있다.
	// (Go는 디폴트로 1개의 CPU 를 사용함)
	runtime.GOMAXPROCS(4)

	fmt.Println("---------")

	// goroutine 활용 1
	// 포인터 형으로 WaitGroup 생성 -> 함수에 전달할 때 & 를 붙이지 않아도 된다.
	wait2 := new(sync.WaitGroup)
	wait2.Add(100)
	for i := 0; i < 100; i++ {
		go hello(i, wait2)
	}

	wait2.Wait()
}

func hello(n int, w *sync.WaitGroup) {
	defer w.Done()

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	fmt.Println(n)
}
