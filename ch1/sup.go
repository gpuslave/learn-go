package main

import (
	"fmt"
	"slices"
	// "go/types"
)

const (
	id   = "id"
	name = "name"
)

func main() {
	// basic variables
	{
		var char rune = 'J'
		var test uint = 0
		var x, y = 123.1, "hello"
		test = test - 2

		fmt.Println(fmt.Println(test, string(char), x, y))
	}

	// using const
	{
		const test_1, test_2 int = 2, 3
		const z int = test_1 + test_2
		fmt.Println(z)
	}

	// arrays and slices
	{
		var arr [3]int = [3]int{1, 2, 3}
		fmt.Println(arr)

		var sl []int = []int{10, 20, 30}
		var app []int = []int{49, 69}

		sl = append(sl, app...)
		fmt.Println(cap(sl))

		var mk []int = make([]int, 0, 15)
		mk = append(mk, 1, 2, 3, 4)
		fmt.Println(mk, len(mk), cap(mk))

		clear(mk) // sets all elements to zero val
		fmt.Println(mk, len(mk), cap(mk))

		splice := make([]float64, 5, 10)
		other := splice[:]
		fmt.Println(slices.Equal(splice, other))

		var copyx []int = nil
		var emptyx []int = make([]int, 0)
		fmt.Println(copyx, slices.Equal(copyx, nil), copyx == nil)
		fmt.Println(emptyx, emptyx == nil)

		copyx = append(copyx, 1, 2, 3)

		y := make([]int, 3)
		copy(y, copyx)

		fmt.Println(y, y == nil)

		var mat [][]int = make([][]int, 0)
		mat = append(mat, []int{1, 2, 3})

		fmt.Println(mat)
	}

	{
		var nilMap map[string]int                    // nil map
		var emtMap map[string]int = map[string]int{} // non-nil map

		fmt.Println(nilMap == nil, emtMap == nil)

		var teams = map[string][]string{
			"First":  {"Mark", "Marcus"},
			"Second": {"Kyle", "Kaysie"},
		}

		fmt.Println(teams)

		val, ok := teams["Third"]
		fmt.Println(val, ok)

		if !ok {
			delete(teams, "Second")
		}
		fmt.Println(teams)

		clear(teams)
		fmt.Println(teams)
	}

	{
		type person struct {
			name string
			age  int
			pet  string
		}

		var mypers person = person{age: 30, pet: "Cat"}
		fmt.Println(mypers)
	}
}
