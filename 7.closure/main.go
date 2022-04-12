package main

import "fmt"

func main() {
	// Go 에서 함수는 클로저로서 사용될 수도 있다.
	// 클로저는 함수 바깥에 있는 변수를 참조하는 함수값을 이른다.
	// python 에서의 개념과 같다고 생각한다.

	// next 공간 안에 nextValue 의 i 변수가 지정되었다고 생각
	next := nextValue()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println("----")

	anotherNext := nextValue()
	fmt.Println(anotherNext()) // 다시 시작
	fmt.Println(anotherNext())
	fmt.Println(anotherNext())
}

// 함수의 캡슐화가 가능하다.
// -> return 되는 함수에 직접 접근할 수 없고,
// 본 함수로부터 return 된 값을 받아오기 때문에
// OOP 의 get(), set() 과 동일하게 움직인다.

// 불필요한 전역 변수를 없애고, 변수를 공유할 수 있게 된다.
// -> java 의 static 과 비슷한 느낌인가?
// -> 이를 Capture 라고 부른다.
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
