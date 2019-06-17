package main

import "fmt"

func main() {
	//定义空接口切片
	//arr := make([]interface{},5}
	var arr []interface{}
	arr = append(arr, 1,3.14,'a',"呵呵",true)

	for _,v := range arr{
		//对数据v进行断言,如果是整型就打印
		//data,ok := v.(int)
		//if ok {
		//	fmt.Println(data)
		//}

		//连起来写
		if data,ok := v.(int);ok == true {
			fmt.Printf("整型数据%d\n",data)
		}else if data,ok := v.(float64); ok {
			fmt.Printf("浮点型数据%g\n",data)
		}else if data,ok := v.(byte); ok {
			fmt.Printf("btye类型数据%v\n", data)
		}else if data,ok := v.(string); ok {
			fmt.Printf("字符串类型数据%s\n", data)
		}
	}
}