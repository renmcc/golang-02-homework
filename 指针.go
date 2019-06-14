package main

import "fmt"

func main22() {
	var a int = 10
	var p *int = &a
	fmt.Println(p)

	//声明了一个指针，默认值为nil（空指针为0）指向了内存地址编号为0的空间
	//0-255为系统占用，不允许用户进行读写操作

	//var p1 *int  //这种叫空指针
	//p1 = 0x042058080 //这种叫野指针  指针变量指向了一个未知的空间
	//
	////访问野指针和空指针对应的内存空间都会报错（C语言可以用野指针，但是也容易报错）
	//*p1 = 123
	//fmt.Println(*p1)
}



//指针变量的new操作
func main2233() {
	//为指针变量创建内存空间
	var p2 *int
	p2 = new(int)
	fmt.Println(*p2)
}

//指针作为函数的参数
func swap(a *int, b *int) {
	*a,*b = *b,*a
}

func main44() {
	a:=10
	b:=20
	swap(&a,&b)
	fmt.Println(a,b)
}

//数组指针（数组的指针）
func main6668() {
	var arr [5]int = [5]int{1,2,3,4,5}

	//定义指针指向数组
	var p5 *[5]int
	p5 = &arr
	fmt.Println(*p5)
	//通过指针获取元素值,需要小括号改变执行顺序
	fmt.Println((*p5)[0])
	//go语言中，对指针优化，可以通过指针操作数组
	fmt.Println(p5[0])
}

//切片指针
func main7777() {
	var slice2 []int = []int{1,2,3,4,5}

	var p6 *[]int
	//因为切片本身就是地址引用，所以此时p6是二级指针
	p6 = &slice2
	fmt.Printf("%p\n",*p6)
	fmt.Printf("%p\n",slice2)
	//切片指针不能通过指针操作元素
	fmt.Println((*p6)[0])

	//切片指针操作切片
	*p6 = append(*p6,6,7,8,9,10)
	fmt.Println(slice2)
}


//切片指针作为函数参数
func test6(s []int) {
	s=append(s,6,7,8,9)
}

//接收一个切片指针
func test7(s *[]int) {
	//s是指针，*s是对应的值
	*s = append(*s, 6,7,8,9)
}

func main8888() {
	s:=[]int{1,2,3,4,5}
	test6(s)
	fmt.Println(s)
	//传递指针到函数
	test7(&s)
	fmt.Println(s)

}

//new创建切片指针
func main999() {
	//空指针不能添加元素
	var p *[]int
	fmt.Printf("%p\n", p)


	p = new([]int)
	fmt.Printf("%p\n",p)
	//利用指针添加数据
	*p=append(*p, 1,2,3,4)
	fmt.Println(*p)
}


//指针数组（数组里存的元素是指针）
func main00011() {
	//元素为指针
	var arr [2] *int
	a,b := 10,20
	arr[0] = &a
	arr[1] = &b

	//获取第一个元素值
	fmt.Println(*arr[0])

}

//指针切片
func main55555() {
	var slice [] *int
	a,b,c := 10,20,30
	slice = append(slice, &a,&b,&c)
	d := 11
	slice[0] = &d

	fmt.Println(*slice[0])
}

//结构体指针
type student struct {
	id   int
	name string
	age  int
	sex  string
}

func main444555() {
	//定义结构体变量
	var stu student =student{ 101,"哆啦A梦",100,"男"}

	//定义结构体指针
	var p *student
	p = &stu
	fmt.Println(*p)

	//通过结构体指针修间接改结构体变量
	(*p).name = "大熊"
	fmt.Println(*p)
	//go语言中，可以通过指针直接修改结构体元素
	p.name = "静香"
	p.sex = "女"
	fmt.Println(*p)

}


//结构体切片
func main00000() {
	var stu = make([]student, 1)
	stu = append(stu,student{102,"小猪佩奇",10,"男"})
	fmt.Printf("%v\n", stu)

	//利用指针在切片首元素添加一个结构体元素
	p := &stu
	(*p)[0]=student{103,"野猪佩奇",20,"男"}
    *p = append(*p, student{104,"花猪佩奇",20,"女"})

	fmt.Printf("%d\n",len(*p))
	fmt.Println(*p)
}

//多级指针
func main() {
	a := 10
	p := &a
	pp := &p
	ppp := &pp
	//一级指针存储变量的地址，二级指针存储一级指针的地址，三级指针存储二级指针的地址
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", pp)
	fmt.Printf("%T\n", ppp)
}











