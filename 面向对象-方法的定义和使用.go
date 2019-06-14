package main

import "fmt"

//内置类型不能绑定方法，用type给内置类型起别名
type Int int

//给类型绑定方法
// func (方法接收者) 方法名(参数列表) 返回值列表
func (a Int) add_sum(b Int) Int {
	return a+b
}

func (a Int) add_sum2(b Int, c Int) Int {
	return b+c
}

//定义结构体方法
type Stu struct {
	name string
	age int
	sex string
}
//值传递
func (s Stu) PrintlnInfo() {
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.sex)
}
//地址传递
func (s *Stu) EditInfo(name string, age int, sex string) {
	//因为结构体指针可以直接调用成员，所以不用(*s).name
	s.name = name
	s.age = age
	s.sex = sex
}

func main() {
	var a Int
	a = 10
	value := a.add_sum(20)
	fmt.Println(value)

	value2 := a.add_sum2(20,20)
	fmt.Println(value2)
	//结构体对象
	s := Stu{"刘姥姥",11,"男"}
	s.PrintlnInfo()
	//结构体是值传递，想要创建修改方法直接修改变量，需要传递指针
	(&s).EditInfo("林黛玉",15,"女")
	//go语言优化，结构体指针方法可以接受变量
	s.EditInfo("林黛玉",15,"女")
	s.PrintlnInfo()

	//结构体方法可以直接操作指针
	var s1 *Stu
	s1 = new(Stu)
	s1.EditInfo("张三",22,"男")
	s1.PrintlnInfo()
}