package main

import "fmt"

type persion5 struct {
	name string
	age int
	sex string
}

type student5 struct {
	persion5
	score int
}

func (p *persion5) PrintInfo(){
	fmt.Printf("大家好，我是%s,我今年%d岁，我是%s生\n",p.name,p.age,p.sex)
}

//和父类方法一样，会重写
func (s *student5) PrintInfo() {
	fmt.Printf("大家好，我是%s,我今年%d岁，我是%s生,我的分数是%d分\n",s.name,s.age,s.sex,s.score)
}

func main() {
	s := student5{persion5{"张三",15,"男"},66}
	//调用子类方法,采用就近原则
	s.PrintInfo()
	//调用父类方法
	s.persion5.PrintInfo()
}