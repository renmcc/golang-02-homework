package shuzu

import "fmt"

func maopao() {
	var arr = [10]int{1,3,5,6,1,2,5,12,32,43}

	//外层控制行
	for i:=0;i<len(arr)-1;i++ {
		//内层控制列
		for j:=0;j<len(arr)-1-i;j++{
			//对比相邻数据满足条件数据交换
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}