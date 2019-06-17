package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var  arr [10]int
	//生成随机数种子
	rand.Seed(time.Now().UnixNano())
	//生成10个10-20之间的随机数
	for i:=0; i<10;i++ {
		arr[i]=rand.Intn(11)+10
	}

	fmt.Println(arr)
}