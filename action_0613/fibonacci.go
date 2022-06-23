package main

import (
	"fmt"
)

func FibonacciNumbers() {
	idx, res := 0, 0
	for i := 0; i <= 10; i++ {
		idx, res = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", idx, res)
	}
}

func fibonacci(n int) (idx int,res int) {
	if n <= 1 {
		res = 1
	} else {
		_, n1 := fibonacci(n - 1)
		_, n2 := fibonacci(n - 2)
		res = n1 + n2
	}
	idx = n
	return
}