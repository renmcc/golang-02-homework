package main

import "fmt"

type TestA struct {
	name  string
	id int
}

type TestB struct {
	TestA
	sex string
	age int
}
//结构体不能嵌套本结构体
//结构体可以嵌套本结构体指针
type TestC struct {
	TestB
	//*TestC
	score int
}

type TestD struct {
	name  string
	id int
}

type TestE struct {
	sex string
	age int
}

type TestF struct {
	TestD
	TestE
	score int
}

func main() {
	var s TestC
	s.name = "李四"
	s.id = 201
	s.age = 20
	s.score = 10
	fmt.Println(s)

	s1 := TestF{TestD{"赵武",12},TestE{"男",18},100}
	fmt.Println(s1)
}