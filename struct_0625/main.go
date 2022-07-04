package main

import "fmt"

func main() {
	person := NewPersonStruct()
	person.name = "张三"
	person.age = 35
	person.gender = 1
	person.provinces = "广东"
	person.city = "深圳"
	person.detailed = "池"
	person.work.name = "腾讯"
	person.work.provinces = "广西"
	fmt.Println(person)
}
