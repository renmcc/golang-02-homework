package main

import "fmt"

//子集
type Hunmaner interface {
	SayHi()
}

//超集
type Persioner interface {
	Hunmaner
	Sing(string)
}

//定义结构体
type students13 struct {
	name string
	age int
	sex string
}

//结构体方法
func (s *students13) SayHi() {
	fmt.Printf("大家好，我是%s,我几年%d岁了，我是%s生\n",s.name,s.age,s.sex)
}
//结构体方法
func (s *students13) Sing(name string) {
	fmt.Printf("大家好，我叫%s，我给大家唱一首%s\n",s.name,name)
}

func main() {
	//子集
	var h Hunmaner
	h = &students13{"韩红",33,"女"}
	h.SayHi()

	//超集
	var p Persioner
	p = &students13{"王菲",11,"男"}
	p.SayHi()
	p.Sing("传奇")

	//将超集转成子集
	//子集不可以转成超集
	h=p
	h.SayHi()
}