package main

import "fmt"

func main() {
	//定义并初始化字典x
	x := make(map[string]string)

	x["A"] = "Apple"
	x["B"] = "Banana"
	x["O"] = "Orange"


	//判断不存在的变量
	if val, ok := x["C"]; ok {
		fmt.Println(val)
	}


	for key , val := range x {
		fmt.Println("Key:", key, "Value:", val)
	}
}
