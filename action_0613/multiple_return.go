package main

import "fmt"

// FuncReturn 练习函数匿名返回值与命名返回值
func FuncReturn() {
	a, b, c := returnValAnonymous()
	multiPly3Nums(a, b, c)
	fmt.Println("returnValAnonymous: ", a, b, c)
}

func multiPly3Nums(a int, b int, c int) int {
	//尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
	//如果返回值具有相同类型，可以传递一个切片给函数，或者是传递一个结构体，因为传递一个指针允许直接修改变量的值，消耗也更少。
	return a * b * c
}

func returnValAnonymous() (int,int,int) {
	a, b, c := 1, 2, 3
	//尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
	//如果返回值具有相同类型，可以传递一个切片给函数，或者是传递一个结构体，因为传递一个指针允许直接修改变量的值，消耗也更少。
	return a + b + c, a * b * c, c - b - a
}

func returnVal() (a int,b int,c int) {
	a = 1 + 2 + 3
	b = 1 * 2 * 3
	c = 3 - 2 - 1
	return a, b, c //Or return
}