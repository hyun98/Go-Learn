package main

import "fmt"

func main() {
	nameList := make([]string, 0, 10)
	var name string

	for {
		fmt.Scanln(&name)
		if name == "0" {
			break
		} else {
			nameList = append(nameList, name)
		}
	}

	for _, n := range nameList {
		defer fmt.Println(n)
	}
}