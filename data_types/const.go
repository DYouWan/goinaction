package main

import (
	"fmt"
)

// Pi 隐式类型定义 如果常量类型没有指定，运行时会根据值自动推断值类型
const Pi = 3.1415926
// OrgName 显示类型定义
const OrgName string = "中石化"
// ZhouYi ZhouEr ZhouSan 常量支持并行定义的方式
const ZhouYi,ZhouEr,ZhouSan = 1,2,3
//使用原括号声明 常量
const (
	Nan = 1
	Nv  = 2
)

//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
const  (
	YiNianJi,N = iota,iota
	ErNianJi = iota + 1
	SanNianJi = iota
)

func GetConst() {
	fmt.Println(Pi, OrgName, ZhouYi, ZhouEr, ZhouSan, Nan, Nv)
	fmt.Println(YiNianJi,N, ErNianJi, SanNianJi)
}

