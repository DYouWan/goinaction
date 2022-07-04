package main

import "fmt"

type Myint int

func CustomType() {
	var age Myint
	age = 10
	fmt.Println(age)
	fmt.Printf("type of a:%T\n", age)
}