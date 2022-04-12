package main

import "fmt"

func main() {
	//array()
	slice()
}

func array() {
	// 배열의 기본 값은 0 이다.
	var a [3]int

	for _, i := range a {
		fmt.Println(i)
	}

	// 배열 초기화
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3} // 배열 크기 생략
	fmt.Println(a1, a2)

	// 다차원 배열
	var multiArr [3][4][5]int
	multiArr[0][1][2] = 3
	for _, i := range multiArr {
		fmt.Println(i)
	}

	var mul = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}
	fmt.Println(mul[1][2])
}

func slice() {
	// 슬라이스 변수 선언
	// 배열과 달리 크기는 지정하지 않는다.
	// 차후 그 크기를 동적으로 변경할 수도 있고, 부분 배열을 발췌할 수도 있다.
	var a []int
	a = []int{1, 2, 3}

	a[1] = 10
	fmt.Println(a)

	// Go 의 내장함수 make() 함수를 이용할 수 있다.
	// make() 함수로 슬라이스를 생성하면, 슬라이스의 길이와 용량을 임의로 지정할 수 있다.
	// make(슬라이스 타입, 슬라이스 길이, 슬라이스 최대 용량)
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))

	// 별도의 길이와 용량을 지정하지 않느면, 기본적으로
	var s2 []int
	if s2 == nil {
		fmt.Println("Nil Slice")
	}
	fmt.Println(len(s2), cap(s2))

	// 부분 슬라이스
	// python 의 장점을 이어받음
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s1[2:5])
	fmt.Println(s1[:4])
	fmt.Println(s1[3:])

	// 슬라이스 추가, 병합, 복사
	s3 := []int{0, 1}
	s3 = append(s3, 2)       // 하나 추가
	s3 = append(s3, 3, 4, 5) // 여러 개 추가
	fmt.Println(s3)

	// append() 슬라이스에 데이터를 추가할 때.
	// 슬라이스 용량이 아직 남아있는 경우, 그냥 데이터를 추가한다.
	// 용량을 초과하는 경우, 현재 용량의 2배에 해당하는 새로운 array를 생성하고
	// 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다.
	// -> C++ 에서 vector 의 push_back 과 비슷한데..
	// Go 에서는 기존 배열 값들을 새로운 배열로 복제한다는 부분이 걸린다.

	sliceA := make([]int, 0, 3)
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)

	// 슬라이스 + 슬라이스
	sliceB := []int{16, 17, 18}
	sliceA = append(sliceA, sliceB...)
	// sliceA = append(sliceA, 16,17,18)
	fmt.Println(sliceA)

	// 슬라이스 복제
	// 슬라이스는 python 과 마찬가지로 배열, 클래스, 객체 등의 자료형은
	// ** 참조 자료형이다 **
	// 소스 슬라이스가 갖는 배열의 데이터를 타겟 슬라이스가 갖는 배열로 복제하는 것
	// => 값의 복사가 일어난다.
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	// **얕은 복사
	copy(target, source)

	fmt.Println(source)
	fmt.Println(target)
	source[1] = 10
	// source 의 값을 변경해도 target 에는 영향을 주지 않는다.
	// int 에 영향을 주지는 않는다.
	fmt.Println(source)
	fmt.Println(target)

	// 이 경우엔 얕은 복사가 일어난다.
	test := target
	target[2] = 20
	// target 의 값을 변경하면 test 에 영향을 미친다.
	fmt.Println(test)
	fmt.Println(target)

	mArr := [][]int{{1, 2, 3}, {4, 5, 6}}
	testmArr := make([][]int, len(mArr), cap(mArr)*2)
	copy(testmArr, mArr)
	mArr[0][1] = 20
	// copy 는 얕은 복사라서 testmArr 의 값에 영향을 준다
	// Go 는 python 처럼 deepcopy 함수가 없다..
	fmt.Println(mArr)
	fmt.Println(testmArr)

}
