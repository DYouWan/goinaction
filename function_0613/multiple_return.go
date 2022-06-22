package main

import "fmt"

func Fibonacci() {
	idx, res := 0,0
	for i := 0; i <= 10; i++ {
		idx, res = fibonacci2(i)
		fmt.Printf("fibonacci(%d) is: %d\n", idx, res)
	}
}

func fibonacci(n int) (idx int,res int) {
	if n <= 1 {
		res = 1
	} else {
		_, v1 := fibonacci(n - 1)
		_, v2 := fibonacci(n - 2)
		res = v1 + v2
	}
	idx = n
	return
}

func fibonacci2(n int) (idx int,res int) {
	if n <= 1 {
		res = 1
	} else {
		_, v1 := fibonacci(n - 1)
		_, v2 := fibonacci(n - 2)
		res = v1 + v2
	}
	idx = n
	return idx, res
}