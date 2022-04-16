package main

import (
	"errors"
	"fmt"
	"log"
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

func divide(a float32, b float32) (result float32, err error) {
	if b == 0 {
		return 0, errors.New("0으로 나누지마")
	}
	result = a / b
	return
}

func divide2(a float32, b float32) (result float32, err error) {
	if b == 0 {
		return 0, fmt.Errorf("%.2f으로 나누지마", b)
	}
	result = a / b
	return
}

func main() {
	a := 10
	b := 20
	res, err := isSame(a, b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

	var num1, num2 float32

	fmt.Scanln(&num1, &num2)
	//result, err := divide(num1, num2)
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)

	fmt.Scanln(&num1, &num2)
	result2, err2 := divide(num1, num2)
	if err2 != nil {
		log.Print(err2)
	}
	fmt.Println(result2)
}
