package main

/*
1、编译时异常
2、编辑时异常
3、运行时异常
 */

import (
	"errors"
	"fmt"
)

func test55(a int, b int) (value int, err error) {
	if b == 0{
		err = errors.New("0不能作为除数")
		return
	}else {
		value = a/b
		return
	}
}

func main() {
	ret,err := test55(10,0)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(ret)
	}
}