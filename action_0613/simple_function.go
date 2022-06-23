package main

import "fmt"

// FunctionParameter 练习函数参数的值传递与引用传递
func FunctionParameter()  {
	a:=1
	callByValue(a)
	callByReference(&a)
	fmt.Println("function_parameter: ",a)
}

func callByValue(a int)  {
	//Go 默认使用按值传递来传递参数，也就是传递参数的副本。
	//函数接收参数副本之后，在使用变量的过程中对副本的值进行的更改，不会影响到原来的变量
	a = 10
	fmt.Println("call_by_value: ",a)
}

func callByReference(a *int)  {
	//如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；
	//通过这个指针的值来修改这个值所指向的地址上的值,那么所有指向这个地址的变量的值都会改变；
	//指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址,所以，按引用传递也是按值传递；

	//几乎在任何情况下，传递指针（一个 32 位或者 64 位的值）的消耗都比传递副本来得少。
	//在函数调用时，切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）
	*a = 5
	fmt.Println("callByReference: ",*a)
}
