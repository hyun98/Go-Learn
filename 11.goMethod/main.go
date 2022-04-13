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

// struct 상속(임베딩)
type person struct {
	name string
	age int
}
func (p *person) greeting1() {
	fmt.Println("person!! greet!")
}

type animal struct {
	spec string
}
func (a *animal) greeting2() {
	fmt.Println("animal!! greet!")
}

type Student1 struct {
	p person // 학생 구조체 안에 사람 구조체를 필드로 가지고 있음. Has-a 관계
	school string
	grade  int
}

type Student2 struct {
	person // 필드명 없이 타입만 선언하면 포함(Is-a) 관계가 됨
	animal // 2개도 가능. 하지만 타입의 메서드 명이 같은게 있다면 컴파일 에러 발생
	// 그냥 2개 쓰지말자.
	school string
	grade  int
}

func main() {
	rect := Rect{width: 10, height: 20}
	area := rect.area()
	rect.area2()

	fmt.Println(rect.width)
	fmt.Println(area)

	// struct 임베딩
	s := Student1{}

	// student 는 p 를 가지고 있기 때문에 greeting 호출 가능
	s.p.greeting1()

	s2 := Student2{}
	// student 는 person 을 상속받는다고 가정할 수 있음.
	s2.greeting1()
	s2.greeting2()
}
