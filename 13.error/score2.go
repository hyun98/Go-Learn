package main

import "fmt"

func average() float64{
	var num int
	fmt.Scanln(&num)

	if num <= 0 {
		panic("잘못된 과목 수입니다.")
	}
	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 {
			panic("잘못된 점수입니다.")
		}
		total += score
	}

	avg := float64(total) / float64(num)

	return avg
}


func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

			main()
		}
	}()

	result := average()
	fmt.Println(result)
}