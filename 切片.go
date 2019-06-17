package main

import "fmt"


//如果整体容量没有超过1024字节，容量按倍数增长，超过1024按之前的1/4增长
func main011() {
	s := make([]int, 5)
	s[0]=123
	s[1]=234
	s[2]=345
	s[3]=456
	s[4]=567
	//增加长度和元素，是在末尾增加数据
	s=append(s,567,789)
	fmt.Println(s)
	fmt.Println("长度：",len(s))
	fmt.Println("容量：",cap(s))
	//增加长度和元素
	s=append(s,44,3,2,2,33)
	fmt.Println("长度：",len(s))
	fmt.Println("容量：",cap(s))
}

func main022() {
	var s []int
	s = append(s,1,2,3)
	fmt.Println(s,len(s))
}

func main04443() {
	s1:= []int{1,2,3,4,5,6}
	s2 := make([]int,6)
	s3 := make([]int,4)
	copy(s2,s1)
	copy(s3,s1[1:5])

	s1=append(s1,1,2,3,4,5)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func BubbleSort(s []int) {
	for i:=0;i<len(s)-1;i++{
		for j:=0;j<len(s)-1-i;j++{
			if s[j] > s[j+1] {
				s[j],s[j+1] = s[j+1],s[j]
			}
		}
	}
}

func main() {
	s1 := []int{12,3,4,5,3,2}
	fmt.Println(s1)
	BubbleSort(s1)
	fmt.Printf("%v\n",s1)
}