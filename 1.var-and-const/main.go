package main

import "fmt"

func main() {
	// 전통적인 변수 할당 방식
	// Go 특징 1. 변수를 할당하고 사용하지 않으면 컴파일 에러가 발생한다.
	var a int = 10
	var f float32 = 11.

	a = 15
	f = 12.55

	// 동일한 타입의 변수 복수개 지정가능
	var i, j, k int = 1, 2, 3

	// 변수를 선언하면서 초기값을 지정하지 않으면,
	// Go는 Zero Value를 기본적으로 할당한다.
	// 즉, 숫자형에는 0, bool 타입에는 false,
	// 그리고 string 형에는 "" (빈문자열)을 할당한다.

	// Go는 변수 타입을 자동으로 추론해서 적용해준다.
	// 자동 적용된 변수 타입은 바뀔 수 없다.
	var str = "hello"
	var it = 11
	// str = 11 -> X
	// it = "hi" -> X

	// Go 만의 변수 할당 방식
	gogo := 13

	// 컴파일 에러 fix
	fmt.Println(gogo, str, it, a, f, i, j, k)

	// const 팁
	// 1. 첫 변수에 iota를 사용하면 CPP의 ENUM과 동일하게 사용 가능하다.
	// -> 상수값을 0 부터 순차적으로 부여한다.
	const (
		Apple = iota
		Grape
		Orange
	)
	// 상수는 사용하지 않아도 컴파일 에러가 발생하지 않는다.
	const c int = 10

	fmt.Println(Apple) // 0 출력

	// ** Go 에서 변수명으로 사용할 수 없는 키워드 25개 **
	//break        default      func         interface    select
	//case         defer        go           map          struct
	//chan         else         goto         package      switch
	//const        fallthrough  if           range        type
	//continue     for          import       return       var

}
