package main

import "fmt"

type  animaler interface {
	call(name string) string
}

type dog struct {}
type cat struct {}
func (d dog) call(name string) string {
	return "我的名字叫" + name + " 我会汪汪汪"
}
func (c cat) call(name string) string {
	return "我的名字叫" + name + " 我会喵喵喵"
}

func main() {
	var a animaler
	a = &dog{}
	fmt.Println(a.call("金毛"))
	a = &cat{}
	fmt.Println(a.call("猫"))
}



