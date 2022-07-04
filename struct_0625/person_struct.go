package main

type person struct {
	name   string
	age    int
	gender int
	address
	work
}
type address struct {
	provinces string
	city      string
	detailed  string
}
type work struct {
	name string
	address
}

func NewPersonStruct() *person {
	return &person{}
}