package main

import "fmt"

func main111() {
	//map的长度是自动扩容的
	m := make(map[int]string,1)
	m[1]="a"
	m[10]="11111"
	m[222]="rrfewrwerewrew"

	//第二个值用于判断key是否存在
	v,ok:= m[11]
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("key不存在")
	}

	for i,v:= range m{
		fmt.Println(i,v)
	}
}

func main333() {
	m := map[int]string{10:"aaa",20:"bbbb",30:"cccc",40:"rrrrr"}
	fmt.Println(m)
	//删除元素用delete
	delete(m,1000)
	fmt.Println(m)
}

func daemo(m map[int]string) {
	m[10]="222222222222222222"
}

//map作为参数传递函数是地址传递
func main() {
	m := map[int]string{10:"aaa",20:"bbbb",30:"cccc",40:"rrrrr"}
	fmt.Println(m)
	daemo(m)
	fmt.Println(m)
}
















