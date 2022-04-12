package myPack

import (
	"fmt"
)

// 함수가 대문자로 시작하면
// 외부에서 호출 가능한 public 을 의미한다.

func IntroPublic(d1 func()) {
	introPrivate()
	fmt.Println("myPack Start")

	// 첫 번째 goodbye 실행
	d1()
}

// 소문자로 시작하면 private 으로 취급한다.
func introPrivate() {
	fmt.Println("inner Only func")
}
