package main

import (
	"fmt"
	"math"
	"reflect"
)

// 구조체(struct)가 필드들의 집합체라면,
// interface 는 메서드들의 집합체이다.
// interface 는 타입이 구현해야 하는 메서드 원형들을 정의한다.
// 해당 인터페이스가 갖는 모든 메서드들을 구현하면 된다.

// -> 만약 Rect 가 Shape 을 자신의 인터페이스로 갖고 싶다면
// == Rect 에 implement Shape 를 추가하려면
// Shape 에 정의된 원형을 모두 구현하면 Shape 은 Rect 의 인터페이스가 된다.

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}


// Rect 와 Circle 타입에 대한 Shape 인터페이스 구현
// Shape 인터페이스라는 사실을 따로 명시할 필요 없나..?

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)

	// 빈 인터페이스 (동적 타입)
	// 빈 인터페이스는 interface{} 로 쓸 수 있다.
	// Go의 모든 타입들은 적어도 0개의 메서드를 구현한다.
	// 따라서 빈 인터페이스는 모든 타입을 담을 수 있는 컨테이너로 볼 수 있다.
	// Java 의 모든 객체들이 Object 를 상속받기 때문에 Object 로 일반화가 가능하듯이,
	// C/C++ 에서 void* 와 같이,
	// Go 에서는 빈 인터페이스를 통해서 타입에 관계없이 객체의 전달을 가능하게 해준다.
	var x interface{}
	x = 1
	fmt.Println(reflect.TypeOf(x)) // int
	x = "TOM"
	fmt.Println(reflect.TypeOf(x)) // string
	printIt(x) // TOM

	// Type Assertion
	var a interface{} = 1

	i := a	// i 는 동적 타입
	j := a.(int)
	println(i) //
	fmt.Println(j)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		// 객체지향 특성 사용
		a := s.area() // 인터페이스 메서드 호출
		fmt.Println(a)
	}
}

func printIt(v interface{}) {
	fmt.Println(v)
}
