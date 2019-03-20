package main

import (
	"fmt"
	"sort"
)

func main()  {
	language := map[string]int {
		"java": 0,
		"c":1,
		"python":3,
		"golang":5,
	}

	//声明一个数组切片
	var langes []string
	for k := range language {
		langes = append(langes, k)
	}
	//按字母排序数组
	sort.Strings(langes)

	for _, k := range langes {
		fmt.Println(k, language[k])
	}

}
