package main

import (
	"fmt"
	"os"
)

func main() {
	//path, _ := os.Getwd()
	f, err := os.Open("./14.defer,panic/1.txt")
	defer f.Close() // 3
	if err != nil {

		// RuntimeError 를 발생시킨다.
		panic("파일이 없습니다.")
	}

	// defer 는 지연 호출한 명령의 역순으로 실행된다.
	// main 마지막에 Close 실행
	defer fmt.Println("Line 20") // 2
	defer fmt.Println("Line 21") // 1

	bytes := make([]byte, 1024)
	if _, err = f.Read(bytes); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(bytes))

	panicTest()

	fmt.Println("Hello, world!") // panic이 발생했지만 계속 실행됨

	doNotDivideByZero()
}

// try-catch 처럼 예외를 무시하고 프로그램을 계속 실행시키기 위해서는
// defer 와 recover 를 함께 사용해야 한다.
func panicTest() {
	defer func() {
		r := recover() //복구 및 에러 메시지 초기화
		fmt.Println(r) //에러 메시지 출력
	}()

	var a = [4]int{1, 2, 3, 4}

	for i := 0; i < 10; i++ { //panic 발생
		fmt.Println(a[i])
	}
}

// 재귀함수를 이용해 defer 문에 복구와 동시에 doNot.. 함수를 다시 실행한다.
func doNotDivideByZero() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

			doNotDivideByZero()
		}
	}()

	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	result := num1 / num2
	mod := num1 % num2

	fmt.Println(num1, "/", num2, " = ", result, ", ", mod)
}
