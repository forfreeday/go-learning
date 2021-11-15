package P02interface

import "fmt"

type Computer struct {
}

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

func (phone Phone) Start() {
	fmt.Println("phone..start")
}

func (phone Phone) Stop() {
	fmt.Println("phone..stop")
}

func (Camera Camera) Start() {
	fmt.Println("camera..start")
}

func (Camera Camera) Stop() {
	fmt.Println("camera..stop")
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}
