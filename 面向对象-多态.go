package main

import "fmt"

type persion11 struct {
	name string
	sex string
	age int
}

type student11 struct {
	persion11
	score int
}

type teacher11 struct {
	persion11
	subject string
}

func (s *student11) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁，我的成绩是%d分。\n",s.name,s.sex,s.age,s.score)
}

func (s *teacher11) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁，我教的学科是%s。\n",s.name,s.sex,s.age,s.subject)
}

//接口的实现
type Personer interface {
	SayHello()
}

//多态的实现
//多态是将接口类型作为函数的参数
//多态实现了接口的统一处理
func SayHello(p Personer) {
	p.SayHello()
}

func main() {
	var p Personer
	p = &student11{persion11{"小红","女",18},88}
	SayHello(p)

	p = &teacher11{persion11{"一红","女",22},"盗墓"}
	SayHello(p)

}