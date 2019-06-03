package main

import "fmt"

func main1() {
	b := "hello world"
	fmt.Printf("%s", b)
}

func main() {
	a := "hello world"
	//go语言中一个汉字三个字符，
	b := "传智播客"
	fmt.Println(len(a))
	fmt.Println(len(b))
}