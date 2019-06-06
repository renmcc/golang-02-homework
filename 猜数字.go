package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成随机数种子
	rand.Seed(time.Now().UnixNano())
	//生成100到999之间的随机数，放入数组
	random := make([]int,3)
	random[0] = rand.Intn(9)+1
	random[1] = rand.Intn(10)
	random[1] = rand.Intn(10)

	var num int
	usernum := make([]int, 3)

	var flag int = 0
	for{
		for{
			fmt.Println("请输入一个三位数")
			fmt.Scan(&num)
			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入错误，请重新输入")
		}
		//获取输入的三位用户输入放入数组
		usernum[0] = num/100
		usernum[1] = num/10%10
		usernum[2] = num%10
		for i:=0;i<3;i++{
			if usernum[i] >random[i]{
				fmt.Printf("输入的第%d位数大了\n", i+1)
			}else if usernum[i] < random[i]{
				fmt.Printf("输入的第%d位数小了\n", i+1)
			}else{
				fmt.Printf("恭喜！第%d输入正确\n", i+1)
				flag++
			}
		}
		if flag == 3{
			break
		}else{
			flag=0
		}
	}
	fmt.Printf("随机数：%v",random)
}