package main

/* 给定 slice s[]int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。 */

func Magnify(factor int,s []int) []int {
	if factor == 0 || s == nil {
		return make([]int, factor)
	}
	len := len(s)
	r := make([]int, factor * len)
	if len == 0 {
		return r
	}
	copy(r,s)
	return r
}
