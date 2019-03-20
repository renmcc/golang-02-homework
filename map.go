package main

import "fmt"

func main() {
	//创建并赋值一个字典
	x := map[string] string {
		"A": "Apple",
		"B": "Banana",
		"O": "Orange",
	}

	//删除元素
	if _, ok := x["B"]; ok {
		delete(x, "B")
	}else {
		fmt.Println("key does not exist")
	}

	fmt.Println(x)
}
