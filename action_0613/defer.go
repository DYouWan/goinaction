package main

import "fmt"

func DeferFunction() {
	i, j := 0, 0
	fmt.Println("DeferFunction Start, ", i, j)
	j++
	defer deferClose(i, j)
	i++
	fmt.Println("DeferFunction End, ", i, j)
}

func deferClose(i int,j int) {
	fmt.Println("Defer Close, ", i, j)
}

func DeferCycle() {
	defer fmt.Println("")
	//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}