package main

import "fmt"

func main() {
	var arr2 = [3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}

	for i:=0;i<3;i++{
		for j:=0;j<4;j++{
			fmt.Print(arr2[i][j], " ")
		}
		fmt.Println()
	}
	//二维数组的行数
	fmt.Println(len(arr2))
	//二维数组的列数
	fmt.Println(len(arr2[0]))
}