package main

func IndexFunc(s string, f func(rune) bool, truth bool) int {
	for i, r := range s {
		if f(r) == truth {
			return i
		}
	}
	return -1
}


func IsAscii(c int) bool {
	if c > 255 {
		return false
	}
	return true
}