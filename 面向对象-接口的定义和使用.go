package main

import "fmt"

type persion10 struct {
	name string
	sex string
	age int
}

type student10 struct {
	persion10
	score int
}

type teacher struct {
	persion10
	subject string
}

func (s *student10) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁，我的成绩是%d分。\n",s.name,s.sex,s.age,s.score)
}

func (s *teacher) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁，我教的学科是%s。\n",s.name,s.sex,s.age,s.subject)
}

//定义接口
//接口定义了规则，方法实现了规则
//接口是虚的，方法是实的
//接口中的方法必须有具体的实现
//格式 type 接口名 interface{方法列表}
//方法名(参数列表) (返回值列表)
type Humaner interface {
	//函数的声明，没有具体的实现，具体的实现要根据对象不同，实现方式也不同
	SayHello()
}

func main() {
	stu := student10{persion10{"小明","男",11},66}
	tea := teacher{persion10{"法师 ","女",33},"考古"}
	//stu.SayHello()
	//tea.SayHello()

	//定义接口类型变量
	//接口做了统一的处理 先实现接口 再根据接口实现方法
	//在需求改变时，只需要修改方法
	var h Humaner
	h = &stu
	h.SayHello()

	h = &tea
	h.SayHello()
}