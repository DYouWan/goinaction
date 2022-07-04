package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	a int8
	b int8
	c int8
}
type s2 struct {
	a bool
	b int32
	c byte
	d int64
	e int8
}

func main() {
	t1 := s1{1, 2, 3}
	t2 := s2{false, 30, 'w',360000,9}
	fmt.Printf("t1.a %p\n", &t1.a) //n.a 0xc0000ac058
	fmt.Printf("t1.b %p\n", &t1.b) //n.b 0xc0000ac059
	fmt.Printf("t1.c %p\n", &t1.c) //n.c 0xc0000ac05a

	fmt.Printf("t2.a %p\n", &t2.a) //n.a 0xc0000ac070
	fmt.Printf("t2.b %p\n", &t2.b) //n.b 0xc0000ac074
	fmt.Printf("t2.c %p\n", &t2.c) //n.c 0xc0000ac078

	fmt.Printf("t2 size: %d, align: %d\n", unsafe.Sizeof(t2), unsafe.Alignof(t2))
}