package main

import "fmt"

//iota实现枚举，自增
const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

const (
	abc = iota
	bcd
	efg
)

func main() {
	a, b, c := 1, true, "def"
	fmt.Println(a,b,c)
	fmt.Println(b,kb,mb,tb,pb)
	fmt.Println(abc,bcd,efg)
}
