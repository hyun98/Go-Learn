package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Go 프로그래밍 언어는 다음과 같은 기본적인 데이타 타입들을 갖고 있다.
	//
	//불리언 타입
	//bool
	//문자열 타입
	//string: string은 한번 생성되면 수정될 수 없는 Immutable 타입임
	//정수형 타입
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//Float 및 복소수 타입
	//float32 float64 complex64 complex128
	//기타 타입
	//byte: uint8과 동일하며 바이트 코드에 사용
	//rune: int32과 동일하며 유니코드 코드포인트에 사용한다

	rawStr := `자유로
운 글\n 
백슬래시를 
인식하지 않는다
라인도 자유롭다`

	str := "일반적인 str\n" +
		"자바와 거의 같다"
	fmt.Println(rawStr)
	fmt.Println(str)

	// 타입 변환, 타입 캐스팅
	var i int = 100
	var u uint = uint(100)
	var f float32 = float32(i)
	fmt.Println(f, u)

	// 타입을 확인하고 싶다면 reflect.TypeOf()를 사용한다.
	fmt.Println(reflect.TypeOf(u))
	fmt.Println(reflect.TypeOf(f))

	str2 := "ABC"
	// 바이트로 변환하면 각 문자의 ASCII 코드가 리스트로 반환된다.
	bytes := []byte(str2)
	str3 := string(bytes)
	fmt.Println(bytes, str3)

}
