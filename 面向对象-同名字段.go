package main

import "fmt"

//当父类中有子类中相同字段时如何处理
type Persion7 struct {
	name string
	age int
	sex string
}

type Student7 struct {
	Persion7
	id int
	name string
	score int
}

func main() {
	var stu Student7
	//采用就近原则，使用子类字段
	stu.name = "张三丰"
	//如果要用父类name
	stu.Persion7.name = "张三"
	fmt.Println(stu)

	stu7 := Student7{Persion7{"张四",18,"男"},10,"张四风",100}
	fmt.Println(stu7)
}