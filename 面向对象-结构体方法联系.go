package main

import "fmt"

type student2 struct {
	name string
	age int
	sex string
	cscore int
	mscore int
	escore int
}

//方法1 打招呼
func (s *student2) SayHello() {
	fmt.Printf("大家好，我叫%s,我今年%d岁了，我是%s生。\n",s.name,s.age,s.sex)
}
//方法2 打印总成绩
func (s *student2) PrintlnScore(){
	score_sum := s.cscore + s.mscore + s.escore
	fmt.Printf("总成绩为%d，平均成绩%d\n", score_sum, score_sum/3)
}
//方法3为对象赋值
func (s *student2) InitInfo(name string, age int, sex string, cscore int, mscore int, escore int) {
	s.name = name
	s.age = age
	s.sex =sex
	s.cscore = cscore
	s.mscore = mscore
	s.escore = escore
}

func main() {
	stu := student2{"贾宝玉",22,"男",55,66,33}
	stu.SayHello()
	stu.PrintlnScore()

	var stu2 student2
	stu2.InitInfo("林黛玉",33,"女",22,33,44)
	stu2.SayHello()
	stu2.PrintlnScore()
}