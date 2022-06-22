package main

import "fmt"

type calculation func(int,int)int

func FunctionType() {
	var f1 calculation
	f1 = add
	fmt.Println(f1(10, 20))

	f := add                        		// 将函数add赋值给变量f1
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
