package main

import "fmt"

//父类
type Opt struct{
	num1 int
	num2 int
}
//加法子类
type AddOpt struct {
	Opt
}

func (add *AddOpt) Operate() int {
	return add.num1+add.num2
}

//减法子类
type SubOpt struct {
	Opt
}

func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}

//设计模式 对于面向对象给予M(模型)V(视图)C(控制器)
//工厂
type OptFratory struct {

}

//num1 值1 num2 值2 op 运算符
func (of *OptFratory) OptCalc(num1, num2 int, op string) (value int) {
	//通过运算符进行分类的计算
	switch op {
	case "+":
		add := AddOpt{Opt{num1,num2}}
		value = add.Operate()
	case "-":
		sub := SubOpt{Opt{num1,num2}}
		value = sub.Operate()
	}
	return
}

//定义接口
type Oper interface {
	Operate() int
}

func main() {
	//s := AddOpt{Opt{1,2}}
	//s1 := SubOpt{Opt{7,2}}
	//var Op Oper
	//Op = &s
	//fmt.Println(Op.Operate())
	//Op = &s1
	//fmt.Println(Op.Operate())
	var opt OptFratory
	value := opt.OptCalc(5,2,"+")
	fmt.Println(value)

}