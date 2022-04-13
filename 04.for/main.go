package main

import "fmt"

func main() {
	// 일반적인 for 문
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for 문 - 조건식만 쓰는 for 루프
	// Go 에는 while 문이 없다.
	// 그래서 for 문을 while 문 처럼 사용할 수 있다.
	n := 1
	for n < 100 {
		n *= 2
		if n > 90 {
			break
		}
	}
	fmt.Println(n)

	// for 문 - 무한루프
	cut := 0
	for {
		fmt.Println("Infinite loop")
		if cut++; cut == 100 {
			break
		}
	}
	fmt.Println("cut : ", cut)

	// for range 문
	// for range 문은 컬렉션으로 부터 한 요소(element)씩 가져와
	// c 차례로 for 블럭의 문장들을 실행한다.
	// 이는 다른 언어의 foreach와 비슷한 용법이다.
	// python의 range와 매우 흡사하다.

	names := []string{"A", "B", "C"}
	for idx, name := range names {
		fmt.Println(idx, name)
	}

	// break, continue, goto 문
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	fmt.Println(a)

END:
	fmt.Println("End")

	// break 레이블 사용하기
	i := 0

	// for 문의 레이블을 설정해 주었다.
L1:
	for {
		if i == 0 {
			break L1 // L1 레이블 탈출
		}
	}
L2:
	for {
		fmt.Println("go")
		break L2
	}

	fmt.Println("OK")
}
