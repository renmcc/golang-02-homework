package main

import "fmt"

//去除空元素
func main444333() {
	data := []string{"red","black","","pink","blue"}

	recive_data := make([]string,1)

	for _,v := range data{
		if v != "" {
			recive_data = append(recive_data, v)
		}
	}
	fmt.Println(recive_data)
}

//切片去重
func main() {
	data2 := []int{1,1,1,2,3,4,5,5,5,3,2,2,4,42,4,2}
	revice_data2 := make([]int,1)
	for _,v := range data2{
		flag := 0
		for _,j := range revice_data2{
			if v == j {
				flag++
				break
			}
		}
		if flag == 0 {
			revice_data2 = append(revice_data2, v)
		}
	}
	fmt.Println(revice_data2)
}