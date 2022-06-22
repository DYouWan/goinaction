package main

/*
	函数的主要知识点:
	1.函数参数：  可变参数、引用传递、值传递  simple_function.SimpleFunc()
	2.函数返回值： 匿名或命名多返回值  function_return.Fibonacci()
	3.变量作用域： 全局与局部变量
	4.定义函数类型与函数类型变量 function_type.FunctionType()
	5.将函数作为参数 function_parameter.IndexFunc 将函数作为返回值 function_return.Do
	6.匿名函数与闭包 function_anonymous
	7.defer语句 return_defer
*/
func main() {
	//idx := IndexFunc("测试 阿松大", unicode.IsSpace, true)
	//fmt.Println(idx)

	FunctionDefer()
}
