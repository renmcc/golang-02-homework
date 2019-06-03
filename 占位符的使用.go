package main

import "fmt"

func main001() {
	var a int = 10
	var b float64 = 10
	var c bool = true
	var d byte = 'A'
	var e string = "hello"
	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", b)
	fmt.Printf("%t\n", c)
	fmt.Printf("%c\n", d)
	fmt.Printf("%s\n", e)
	//打印指针
	fmt.Printf("%p\n", &a)
	//打印变量的类型
	fmt.Printf("%T\n", a)
}

func main() {
	a := 123
	b := 0123
	c := 0x123
	//go中不能直接打印二进制
	fmt.Println(a,b,c)
	fmt.Printf("%b\n%b\n%b\n",a,b,c)
}