package main

import "fmt"

type Student struct {
	id      int
	name    string
	sex     string
	age     int
	address string
}

func main() {
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










