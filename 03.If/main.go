package main

import (
	"fmt"
)

func main() {
	// if 문은 괄호()가 필요없다
	// else, else if 문은 이전 if 가 종료되는 중괄호 바로 옆에 있어야 한다.
	// 그리고 반드시 중괄호로 감싸서 블럭을 지정해 주어야 한다.
	k := 1
	if k == 1 {
		fmt.Println("ONE")
	} else {
		fmt.Println("no")
	}

	// 조건식을 사용하기 이전에 간단한 문장을 함께 실행할 수 있다.
	i := 70
	if v := i * 2; v < 150 {
		fmt.Println("less then 150")
	}
	// v+= 10 -> scope를 벗어난 에러

	// Go 의 switch 는 break 가 필요없다.
	// switch 문 뒤에 변수 또는 표현식을 작성할 수 있다.
	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	fmt.Println("basic switch : " + name)

	// Expression을 사용한 경우
	switch x := category << 2; x - 1 {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	fmt.Println("expression switch : " + name)

	// Go 특징 : switch 뒤에 표현식이나 변수가 없어도 된다.
	// 복잡한 elseif 문을 단순하게 표현할 수 있다.
	category = 2
	switch {
	case category == 1:
		name = "Paper Book"
	case category == 2:
		name = "eBook"
		fallthrough // Go는 기본적으로 break 가 걸려있는데, break 를 무시하고 내려가게 하는 명령어
	case category == 3 || category == 4:
		name = "Blog"
		fallthrough
	default:
		name = "default 도달"
	}
	fmt.Println("no value or expression switch and fallthrough : " + name)

	// Go 특징 : 타입을 통해 switch 문을 검사할 수 있다.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
