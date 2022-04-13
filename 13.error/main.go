package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Go 에러는 error 인터페이스를 통해서 주고 받는다.
// interface 는 Error() 메서드 하나를 갖는다.
// 이 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있다.

type myError struct {
	Code string
	Message string
}

func (e *myError) Error() string {
	return e.Code + ", " + e.Message
}

func do() error{
	return &myError{Code:"400", Message:"bad request"}
}

func do2() error{
	return errors.New("error!!")
}


func main() {
	f, err := os.Open("./temp.txt")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(f.Name())
	}

	er1 := do()
	fmt.Println(er1)
	er2 := do2()
	fmt.Println(er2)
}
