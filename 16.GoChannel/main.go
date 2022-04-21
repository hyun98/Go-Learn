package main

import "fmt"

// 채널은 make() 함수를 통해 미리 생성되어야 하며, 채널 연산자 <- 을 통해 데이터를 주고받는다.
// 보통 goroutine 들 사이에 데이터를 주고 받는데 사용된다.
// 별도의 lock 을 걸 필요없이 상대방과 데이터를 동기화 할 수 있다,
func main() {
	ch := make(chan int)

	// 고루틴 생성
	go func() {
		ch <- 123	//  채널에 123을 보낸다.
	}()

	var i int
	i = <- ch // 채널로부터 123을 받는다.
	fmt.Println(i)

	// Go 채널은 수신자와 송신자가 서로를 기다리는 속성을 가지고 있다.
	// 이를 활용해서 Go 루틴이 끝날 때까지 기다리는 기능을 구현할 수 있다.
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	bufferingTest()
	// 위의 Go 루틴이 끝날 때까지 대기
	<- done // 주석 처리를 하면 1~9 출력이 되지 않고 main 함수가 바로 return 된다.

	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

}

func bufferingTest() {
	// Go 채널은 Unbuffered Channel 과 Buffered Channel 이 있다.
	// Unbuffered Channel 은 하나의 수신자가 데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여있게 된다.

	c := make(chan int, 1)
	c <- 1
	var i int
	i = <- c
	fmt.Println(i) // 채널을 받는 수신자가 없다.
}

// 송수신 채널
// 송신 파라미터는 (p chan<- int)와 같이 chan<- 을 사용하고,
// 수신 파라미터는 (p <-chan int)와 같이 <-chan 을 사용한다.
// 만약 송신 채널 파라미터에서 수신을 한다거나, 수신 채널에 송신을 하게되면, 에러가 발생한다.
func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 에러발생
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
