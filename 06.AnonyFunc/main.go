package main

import "fmt"

func main() {

	// 람다 기본 사용법
	sum := func(n ...int) int { // 익명 함수 -> 람다
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)

	// Go 에서 함수는 일급함수로 취급된다.
	// 따라서 Go의 기본 타입과 동일하게 취급되며,
	// 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로도 사용될 수 있다.
	// 즉, 함수의 입력 파라미터나 리턴 파라미터로서 함수 자체가 사용될 수 있다.

	add := func(i int, j int) int {
		return i + j
	}

	// 함수 넘겨주기
	r1 := calc(add, 10, 20)
	fmt.Println(r1)

	// 함수를 매개변수 부분에서 바로 생성해서 넘겨주기
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	fmt.Println(r2)
}

// 함수를 인자로 받는 함수
func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// type 은 구조체(struct), 인터페이스 또는 Custom Type 을 정의한다.
// int, int 를 전달받아서 int 를 반환하는 함수를 calculator 라고 정의했다.
type calculator func(int, int) int

func calc2(f calculator, a int, b int) int {
	return f(a, b)
}
