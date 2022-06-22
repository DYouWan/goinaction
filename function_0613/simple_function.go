package main

import "fmt"

func SimpleFunc(){
	a:=1
	callByVal(a)
	callByReferences(&a)

	f1("Q","L","T")
}

func callByVal(a int){
	a = 10
	fmt.Println("callByVal " , a)
}

func callByReferences(a *int){
	*a = 20
	fmt.Println("callByReferences ", *a)
}

func f1(array ...string){
	f2(array)
}
func f2(array []string)  {
	f3(array...)
}
func f3(array ...string){
	fmt.Println(array)
}