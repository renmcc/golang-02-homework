package main

import "fmt"

type Opt struct{
	num1 int
	num2 int
}

type AddOpt struct {
	Opt
}

type SubOpt struct {
	Opt
}

func (add *AddOpt) Operate() int {
	return add.num1+add.num2
}

func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}

func main() {
	s := AddOpt{Opt{1,2}}
	value := s.Operate()
	fmt.Println(value)

	s1 := SubOpt{Opt{5,2}}
	fmt.Println(s1.Operate())
}