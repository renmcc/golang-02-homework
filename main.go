package main

import (
	"fmt"
	"github.com/renmcc/golang-02-homework/qiepian"
)

func main() {
	var arr = []int{1,3,5,6,1,2,5,12,32,43}


	fanzhuan := qiepian.Reverse(arr)
	fmt.Printf("%v\n",fanzhuan)

	paixu := qiepian.Sort(arr,true)
	fmt.Printf("%v\n",paixu)
}



