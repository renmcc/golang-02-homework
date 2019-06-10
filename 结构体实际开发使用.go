package main

import "fmt"

//定义结构体存储5名学生 三门成绩 求出每名学生的总成绩和平均成绩
type students struct {
	id int
	name string
	score []int
}

func main() {
	arr := []students{{10,"张三",[]int{100,77,88}},
		{10,"张四",[]int{100,78,87}},
		{11,"张五",[]int{100,79,83}},
		{12,"张六",[]int{100,73,82}},
		{13,"张七",[]int{100,76,81}},
		}

	for _,v := range arr{
		var SumScore int
		var AverageScore int

		for i:=0;i<len(v.score);i++{
			SumScore += v.score[i]
		}
		AverageScore = SumScore / len(v.score)
		fmt.Printf("%s的总成绩%d平均成绩是%d\n",v.name,SumScore,AverageScore)
	}
}