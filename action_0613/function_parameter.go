package main

import (
	"fmt"
	"unicode"
)

func IndexFunc() {
	i := indexFunc("你好啊张三A", IsAscii, true)
	fmt.Println("字符串中首个Ascii下标为: ", i)
}

func IsAscii(s rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, s)
}

func indexFunc(s string, f func(rune) bool, truth bool) int {
	for i, r := range s {
		if f(r) == truth {
			return i
		}
	}
	return -1
}