package shuzu

import "fmt"

func fanzhuan() {
	var change_arr = [10]int{10,29,33,44,11,110,22,44,88,55}
	start := 0
	end := len(change_arr) - 1
	for {
		if start > end {
			break
		}
		change_arr[start],change_arr[end] = change_arr[end],change_arr[start]
		start++
		end--
	}
	fmt.Println(change_arr)
}