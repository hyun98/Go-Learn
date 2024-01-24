package main

import (
	"CloudNative/Core"
	"context"
	"errors"
	"fmt"
	"time"
)

var count int

func EmulateTransientError(ctx context.Context) (string, error) {
	count++

	if count <= 3 {
		return "intentional fail", errors.New("error")
	} else {
		return "success", nil
	}
}

func main() {
	//ctx := context.Background()
	//serviceCircuit := func(ctx context.Context) (string, error) {
	//	return "success", nil
	//}

	//wrapped := Core.Breaker(Core.DebounceLast(serviceCircuit, time.Millisecond*100), 10)

	//r := Core.Retry(EmulateTransientError, 5, 2*time.Second)
	//
	//res, err := r(context.Background())
	//fmt.Println(res, err)

	//ctxt, cancel := context.WithTimeout(ctx, time.Second*1)
	//defer cancel()
	//
	//timeout := Core.Timeout(Slow)
	//res, err := timeout(ctxt, "hello, world")
	//
	//fmt.Println(res, err)

	/* Funnel 실행 */
	sources := make([]<-chan int, 0) // 빈 수신 채널 슬라이스 생성

	for i := 0; i < 3; i++ {
		ch := make(chan int)
		sources = append(sources, ch)

		go func() {
			defer close(ch)

			for i := 1; i <= 5; i++ {
				ch <- i
				time.Sleep(time.Second)
			}
		}()
	}

	dest := Core.Funnel(sources...)
	for d := range dest {
		fmt.Println(d)
	}
}

func Slow(str string) (string, error) {
	time.Sleep(2 * time.Second)
	return str, nil
}
