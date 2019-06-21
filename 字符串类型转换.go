package main

import (
	"fmt"
	"strconv"
)

func main() {
	//将布尔类型转为字符串
	s := strconv.FormatBool(false)
	fmt.Println(s)

	//将整型转为字符串，第二个参数是进制
	s1 := strconv.FormatInt(123,10)
	fmt.Println(s1)

	//将整型转成字符串10进制
	s2 := strconv.Itoa(567890)
	fmt.Println(s2)

	//将浮点型转成字符串  数据 格式 保留小数位数（四舍五入） 数据格式
	s3 := strconv.FormatFloat(3.141592653,'f',5,64)
	fmt.Println(s3)


	//将字符串类型转成布尔类型 返回2个值
	b,_ := strconv.ParseBool("false")
	fmt.Println(b)

	//将字符串转成整型,返回2个值
	ret,_ := strconv.ParseInt("44444",10,64)
	fmt.Println(ret)

	//将字符串转成整型 10进制
	ret2,_ := strconv.Atoi("55555555")
	fmt.Println(ret2)

	//将字符串转成浮点型
	ret3,_ := strconv.ParseFloat("1.2345563443",64)
	fmt.Println(ret3)


	b3 := make([]byte,0,1024)
	//将bool类型放入指定切片中
	b3 = strconv.AppendBool(b3,false)

	//将整型放入指定切片 切片 整型 进制
	b3 = strconv.AppendInt(b3, 66666,10)

	//将浮点型数据放入指定切   切片 指定数据 格式 小数点后面位数  
	b3 = strconv.AppendFloat(b3,1.2345,'f',5,64)
	fmt.Println(b3,string(b3))
}