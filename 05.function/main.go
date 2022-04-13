package main

import "fmt"

func main() {
	msg := "hello"
	say1(&msg)
	fmt.Println(msg)
	say2(5, "hi", "hello", "world")

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum2(1, 2, 3, 4, 5))
}

// 이 부분은 C 언어의 규칙을 그대로 따르는 모습이다.
// pass by reference
func say1(msg *string) {
	fmt.Println(*msg)
	*msg = "Changed"
}

// 매개변수 여러개 -> 가변 인자 함수
// python 의 *args 와 비슷
func say2(i int, msg ...string) {
	fmt.Println(msg)
	for _, s := range msg {
		fmt.Println(s)
	}
}

// return 타입 지정 #1
func sum(nums ...int) int {
	s := 0
	for _, i := range nums {
		s += i
	}
	return s
}

// return 타입 및 리턴 변수 지정 #2
func sum2(nums ...int) (count int, total int) {
	// count, total 에는 기본 값으로 0이 들어간다.
	// string 이면 ""
	// bool 이면 false
	for _, i := range nums {
		total += i
	}

	count = len(nums)

	// 리턴 변수를 지정해 주었다면 return 문이 반드시 필요하다.
	// (그냥 return 만 적으면 됨)
	return
}
