package main

import (
	"fmt"
	// "go/types"
)

const (
	id   = "id"
	name = "name"
)

func main() {
	var char rune = 'J'
	var test uint = 0
	var x, y = 123.1, "hello"
	test = test - 2

	fmt.Println(fmt.Println(test, string(char), x, y))

	const test_1, test_2 int = 2, 3
	const z int = test_1 + test_2
	fmt.Println(z)

	// arr
	// var arr [3]int = [3]int{1, 2, 3}

	var sl []int = []int{10, 20, 30}
	var app []int = []int{49, 69}

	sl = append(sl, app...)
	fmt.Println(sl)
}
