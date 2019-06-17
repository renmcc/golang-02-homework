package main

import "fmt"

func test8() {
	fmt.Println("hello world")
}

func test9() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}

//defer调用的函数并没有直接执行，而是先加载到栈区内存中，在函数结束时进行调用，从后向前调用
func main99955() {
	//defer fmt.Println("abc")
	//fmt.Println("efg")
	//defer  fmt.Println(123)
	//fmt.Println("呵呵")

	//函数中有返回值不能使用defer
	defer test8()
	test9()
}

func main() {
	a := 10
	b := 20

	//定义函数类型变量f
	//f := func (a int, b int){
	//	fmt.Println(a)
	//	fmt.Println(b)
	//}
	//f(a,b)

	defer func () {
		fmt.Println(a)
		fmt.Println(b)
	}()

	a = 100
	b = 200

	fmt.Println(a,b)

}