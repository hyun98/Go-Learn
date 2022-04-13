package main

import "fmt"

// Go 언어는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 정의된다.
type Rect struct {
	width, height int
}

// Go 메서드는 특별한 형태의 func 함수이다.
// func 키워드와 함수명 사이에 "그 함수가 어떤 struct 를 위한 메서드인지"를 표시한다.
// 이 부분은 Receiver 라고 하며 struct 변수명은 함수 내에서 입력 파라미터처럼 사용된다.
func (r Rect) area() int {
	return r.width * r.height
}

// 포인터 receiver
// 객체의 포인터로 받는다.
func (r *Rect) area2() int{
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{width: 10, height: 20}
	area := rect.area()
	rect.area2()

	fmt.Println(rect.width)
	fmt.Println(area)


}
