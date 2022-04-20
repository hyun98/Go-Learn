package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}

// map 형태의 data 필드를 가지는 mapStruct 를 정의함
type mapStruct struct {
	data map[int]string
}

// 임의의 생성자 함수
func newStruct() *mapStruct {
	d := mapStruct{}          // 구조체 객체를 생성하고 초기화함
	d.data = map[int]string{} // data 필드의 맵을 초기화함
	return &d
}

func main() {
	// person 객체 생성 ver1
	p := person{}

	p.name = "Lee"
	p.age = 10

	// person 객체 생성 ver2
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Son", age: 50}
	fmt.Println(p1, p2)

	// person 객체 생성 ver3
	// new 를 사용하면 모든 필드를 Zero value 로 초기화 하고,
	// person 객체의 포인터 (*person) 를 리턴한다.
	p3 := new(person)
	// p는 포인터지만 '->' 기호가 아닌 '.' 을 사용한다.
	// ** 용법이 포인터가 아닌 객체들과 같지만 타입은 포인트임을 확실히 인지하고 있어야 한다. **
	p3.name = "Lee"
	fmt.Println(p3) // &{Lee 0}

	// 생성자 (constructor)
	// struct 객체를 만들 때
	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(dic) // &{map[1:A]}

}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
