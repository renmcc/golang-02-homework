package main

import "fmt"

type Student struct {
	id      int
	name    string
	sex     string
	age     int
	address string
}

func main222() {
	//通过结构体名定义结构体变量
	//var s Student
	//s.id = 200
	//s.name = "哈哈哈"
	//s.sex = "男"
	//s.age = 12
	//s.address = "北京"
	s := Student{id:101,name:"关羽",sex:"男",age:18,address:"燕郊"}
	s1 := s
	s1.age=55

	//结构体比较
	if s1 == s{
		fmt.Println("相同")
	}else{
		fmt.Println("不相同")
	}
	//结构体成员比较
	if s1.age == s.age {
		fmt.Println("年龄相同")
	}else{
		fmt.Println("年龄不相同")
	}

	//定义结构体数组
	var arr [5]Student
	for i:=0;i<len(arr);i++{
		fmt.Scan(&arr[i].id,&arr[i].name,&arr[i].age,&arr[i].sex,&arr[i].address)
	}

	//结构体排序，根据结构体成员排序
	for i:=0;i<len(arr)-1;i++{
		for j:=0;i<len(arr)-1-j;j++{
			if arr[j].age > arr[j+1].age {
				arr[j].age,arr[j+1].age = arr[j+1].age,arr[j].age
			}
		}
	}

	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

}

func main3334() {
	//结构体切片
	arr := []Student{{1,"曹操","男",33,"山东"},
		{2,"刘备","男",44,"成都"},
		{3,"孙权","男",22,"江南"}}

	//结构体切片添加数据
	arr = append(arr,Student{4,"张飞","男",55,"河北"})

	for i,v := range arr{
		fmt.Printf("%d,%v\n",i,v)
	}
}

//结构体字典
func main444() {
	//定义map
	m := make(map[int]Student)

	m[10] = Student{1,"曹操","男",33,"山东"}
	m[11] = Student{2,"刘备","男",44,"成都"}
	m[12] = Student{3,"孙权","男",22,"江南"}

	for k,v := range m{
		fmt.Println(k,v)
	}
}
//结构体切片字典
func main5555() {
	m := make(map[int][]Student)
	m[10] = []Student{{1,"曹操","男",33,"山东"},
		{2,"刘备","男",44,"成都"},
		{3,"孙权","男",22,"江南"}}
	m[11] = []Student{{1,"曹操","男",33,"山东"},
		{2,"刘备","男",44,"成都"},
		{3,"孙权","男",22,"江南"}}

	for k,v := range m {
		fmt.Println(k,v)
	}
}

//结构体作为函数参数
type person struct {
	id    int
	name  string
	score int
	sex   string
}

func testfunc(s person) {
	fmt.Println(s.id)
	fmt.Println(s.name)
	fmt.Println(s.score)
	s.name="李逵"
}

func main666() {
	stu :=person{101,"宋江",10,"男"}
	testfunc(stu)
	//结构体作为参数传递给函数是值传递
	fmt.Println(stu.name)
}
//结构体切片是地址传递
func testfunc2(s []person) {
	fmt.Println(s[0].id)
	fmt.Println(s[0].name)
	fmt.Println(s[0].score)
	s[0].name="李逵"
}

func main() {
	stus := []person{{101,"宋江",10,"男"},{102,"宋江",10,"男"}}
	stus = append(stus,person{103,"宋江",10,"男"})
	testfunc2(stus)
	fmt.Println(stus)
}


