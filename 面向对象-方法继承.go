package main

import "fmt"

type person3 struct {
	id   int
	name string
	age  int
	sex  string
}

type student3 struct {
	//p     person3 //结构体变量 结构体类型
	person3
	score int
}

//定义父类方法
func (obj *person3) SayHello() {
	fmt.Println("大家好，我是",obj.name,"我是",obj.sex,"生我今年",obj.age)
}

func main() {
	stu := student3{person3{100,"贾玲",18,"女"},100}
	//子类结构体继承父类结构体，允许使用父类的成员和方法
	stu.SayHello()
}
