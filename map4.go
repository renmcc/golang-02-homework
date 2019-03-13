package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["A"] = 10
	x["B"] = 12

	fmt.Println(x)

	//删除一个元素
	delete(x, "B")

	fmt.Println(x)
}
