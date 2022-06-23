package main

import "fmt"

func ClosureFunction(){
	a := func(num int) {
		num++
		fmt.Println(num)
	}
	for i := 0; i < 10; i++ {
		a(i)
	}
}