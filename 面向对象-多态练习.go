package main

import "fmt"

//第一步：接口实现
type USBer interface {
	Read()
	Write()
}

//第二步：创建对象
type USBDev struct {
	id int
	name string
	speed int

}

type Mobile struct {
	USBDev
	call string
}

type Upan struct {
	USBDev
}
//第三步：实现方法
func (m *Mobile) Read() {
	fmt.Printf("%s正在读取数据速度为%d\n",m.name, m.speed)
}

func (m *Mobile) Write() {
	fmt.Printf("%s正在写入数据速度位%d\n",m.name,m.speed)
}

func (u *Upan) Read() {
	fmt.Printf("%s正在读取数据速度为%d\n",u.name, u.speed)
}

func (u *Upan) Write() {
	fmt.Printf("%s正在写入数据速度位%d\n",u.name,u.speed)
}

//第四步：多态的实现
func UseDev (usb USBer) {
	usb.Read()
	usb.Write()
}

func main() {
	//接口类型
	var usb USBer
	usb = &Mobile{USBDev{1,"诺基亚",10},"隔壁老王"}
	UseDev(usb)

	usb = &Upan{USBDev{2,"金士顿",40}}
	UseDev(usb)
}