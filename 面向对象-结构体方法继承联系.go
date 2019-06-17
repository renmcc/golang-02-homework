package main

import "fmt"

/*
记者：我是记者 我的爱好是偷拍 我的年龄是34 我是一个男狗仔
程序员：我叫孙权 我的年龄23 我是男生 我的工作年限是3年
 */

type Person struct {
	name string
	age int
	sex string
}
//记者
type Rep struct {
	Person
	hobby string
}
//程序员
type Dep struct {
	Person
	workertime int
}

func (p Person) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁。", p.name,p.sex,p.age)
}

func (r Rep) SayHi() {
	r.Person.SayHello()
	fmt.Printf("我的爱好是%s\n",r.hobby)
}

func (d Dep) SayHi() {
	d.Person.SayHello()
	fmt.Printf("我的工作年限是%d\n",d.workertime)
}

func main() {
	r := Rep{Person{"卓伟", 40, "男"}, "偷拍"}
	d := Dep{Person{"汤姆逊", 68, "男"}, 40}
	r.SayHello()
	r.SayHi()
	d.SayHello()
	d.SayHi()
}