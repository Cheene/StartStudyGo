package main

import "fmt"

type x struct {
	a int8 // b nit = 1 Byte
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(23),
		b: int8(43),
		c: int8(12),
	}

	fmt.Printf("%p \t %p\t %p\n", &m.a, &m.b, &m.c)
}
