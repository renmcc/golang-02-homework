package main

import "fmt"

func main00() {
	const A = 10
	b := 20
	c := A+b
	fmt.Println(c)
	//10是硬常量
	d := c-10
	fmt.Println(d)
}

func main() {
	const(
		a = iota
		b,c = iota, iota
		d,e
		f = iota
	)
	fmt.Println(a,b,c,d,e,f)
}