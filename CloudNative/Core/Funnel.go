package Core

import "sync"

// Funnel
// N개의 read-only 채널을 인자로 받는다.
func Funnel(sources ...<-chan int) <-chan int {
	dest := make(chan int) // 공유 출력 채널 선언

	// 모든 sources 의 채널이 닫혔을 때 출력 채널을 자동으로 닫기 위해 사용한다.
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(sources)) // waitGroup의 크기를 지정한다.

	// 각 입력 채널에 대해 고루틴을 시작한다.
	for _, ch := range sources {
		go func(c <-chan int) {
			defer waitGroup.Done() // 채널이 닫히면 WaitGroup으로 알려준다.

			for n := range c { // 채널이 닫힐 때까지 계속해서 값을 받아온다.
				dest <- n
			}
		}(ch)
	}

	go func() {
		waitGroup.Wait()
		close(dest)
	}()

	return dest
}
