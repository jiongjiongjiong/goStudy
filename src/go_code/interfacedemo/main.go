package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

func (p Phone) Call() {
	fmt.Println("手机在打电话")
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	c := Computer{}
	p := Phone{}
	ca := Camera{}

	c.Working(p)
	c.Working(ca)
	fmt.Println("********************************")
	fmt.Println("********************************")
	fmt.Println("********************************")
	var usbArr [3]Usb
	usbArr[0] = Phone{"iphone"}
	usbArr[1] = Phone{"mi"}
	usbArr[2] = Camera{"nikon"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
