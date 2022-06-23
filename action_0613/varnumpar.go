package main

import "fmt"

func VariableArityMethods()  {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func F1(s ...string) {
	F2(s...)
	F3(s)
}
func F2(s ...string) { }
func F3(s []string) { }
