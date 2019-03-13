package main

import "fmt"

func main() {
	//定义字典x
	var x map[string]string
	//make内置方法初始化x后才可以赋值
	x = make(map[string]string)

	x["A"] = "Apple"
	x["B"] = "Banana"
	x["O"] = "Orange"


	for key , val := range x {
		fmt.Println("Key:", key, "Value:", val)
	}
}
