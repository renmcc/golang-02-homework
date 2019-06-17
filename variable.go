package main

import (
	"fmt"
	"os/exec"
)

func main01() {
	a := 5
	b := 3.141592
	fmt.Print("a")
	fmt.Print("b")
	fmt.Printf("==%05d==", a)
	fmt.Printf("%.4f", b)
}

func main02() {
	var a int
	var b string
	fmt.Scanf("%3d", &a)
	fmt.Scanf("%s", &b)
	fmt.Println(a,b)
}

func main03() {
	var name string
	var password string
	fmt.Println("请输入账号")
	fmt.Scanf("%s", &name)
	fmt.Println("请输入密码")
	fmt.Scanf("%s", &password)
	fmt.Println(name, password)
}

func main11() {
	var a string = "呵呵"
	var b string = "哈哈"
	fmt.Println(a+b)
}



func main() {
	arg := "curl -I \"www.baidu.com\""
	cmd := exec.Command("/bin/sh", "-c", arg)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}