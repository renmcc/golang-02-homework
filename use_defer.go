package main

import "fmt"

func first() {
	fmt.Println("first func run")
}

func second() {
	fmt.Println("second func run")
}

func main() {
	//defer在main函数执行结束后才会执行
	defer first()
	second()
}
