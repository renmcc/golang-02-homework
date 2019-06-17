package main

import "fmt"

func main1() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num <0 {
		fmt.Println(num, "is negative")
	}else if num < 10 {
		fmt.Println(num, "has 1 digit")
	}else {
		fmt.Println(num, "has nultiple digits")
	}

	b := 2
	switch b {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("null value")
	}

	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("len:", len(a))

	//b := [5]int{1,2,3,4,5}

	var twoD [2][3]int
	
	for i = 0; i < 2; i++ {
		for j := 0; j <3; j++ {
			twoD[i][j] = i+j
		}
	}
}


func main2() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))
	
}

func main3() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	
	fmt.Println("map:", m)
	fmt.Println(m["k1"])
}


func main4() {
	nums := []int{1,2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}


	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s ->%s\n", k,v)
	}

	for i, c := range "golang" {
		fmt.Printf("%d,%s\n",i, c)
	}


}


func plus(a int, b int) int {
	return a + b
}

func main5() {
	res := plus(1,2)
	fmt.Println("1+2=", res)
}


func vals() (int, int) {
	return 3, 7
}

func main6() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
}


func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main7() {
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4,6,888}
	sum(nums...)
}


func main() {
	for i:=0;i<=100;i+=2{
		fmt.Println(i)
	}
}





















