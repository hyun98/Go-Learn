package main

import (
	"fmt"
)

// Go 에러는 error 인터페이스를 통해서 주고 받는다.
// interface 는 Error() 메서드 하나를 갖는다.
// 이 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있다.

// 문자열을 반환해주는 Error() 메소드를 갖는 객체들은
// 모두 에러 타입이 될 수 있다.

type CustomError struct {
	Code    string
	Message string
}

func (e *CustomError) Error() string {
	return e.Code + ", " + e.Message
}

func isSame(a int, b int) (bool, error) {
	if a != b {
		return false, &CustomError{Code: "400", Message: "Not same"}
	}
	return true, nil
}

func main() {
	//f, err := os.Open("./temp.txt")
	//if err != nil {
	//	log.Fatal(err.Error())
	//} else {
	//	fmt.Println(f.Name())
	//}

	a := 10
	b := 20
	res, err := isSame(a, b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
