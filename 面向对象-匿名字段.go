package main

import "fmt"

type persion struct {
	name string
	age int
	sex string
}

//结构体嵌套结构体
type studentss struct {
	//通过匿名字段实现继承
	id int
	persion  //结构体的名称作为另一个结构体的成员
	score int
}

func main() {
	var stu5 studentss
	stu5.id = 1
	stu5.name = "张三"
	stu5.age = 13
	stu5.sex = "男"
	stu5.score = 66
	fmt.Println(stu5)

	stu6 := studentss{2,persion{"李四",15,"男"},77}
	fmt.Println(stu6)
}