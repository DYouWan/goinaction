package main

import "fmt"

func FunctionDefer() {
	x := 1
	y := 2
	defer calcDefer("AA", x, calcDefer("A", x, y))
	x = 10
	defer calcDefer("BB", x, calcDefer("B", x, y))
	y = 20
}

func calcDefer(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

