package main

func test56() {
	//可以在程序中直接调用panic，调用后程序终止执行
	//print和panic是对错误信息进行处理的 不建议程序中运行
	panic("hello world")
}

func main() {
	test56()

}