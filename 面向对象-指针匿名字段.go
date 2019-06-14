package main

import "fmt"

type person8 struct {
	name string
	age int
	sex string
}

type student8 struct {
	//指针匿名字段
	*person8
	id int
	name string
	score int
}

func main() {
	var stu student8
	//指针需要开辟空间，默认指针空间是nil
	stu.person8 = new(person8)
	stu.name = "郭襄"
	stu.person8.name = "郭芙"
	stu.id = 10
	stu.score = 100
	fmt.Println(stu)
	fmt.Println(stu.name)
	fmt.Println(stu.person8.name)

	//直接赋值
	stu8 := student8{&person8{"张三",88,"男"},11,"李四",88}
	fmt.Println(stu8)
	fmt.Println(stu8.name)
	fmt.Println(stu8.person8.name)
}



















