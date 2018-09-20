package main

import "fmt"

type COLOR int

const (
	RED    COLOR = 1
	WHITE  COLOR = 2
	YELLOW COLOR = 4
)

type PhoneDate struct {
	name  string
	price int
	color COLOR
}

func (phone PhoneDate) getname() string {
	return phone.name
}

func (phone PhoneDate) getprice() int {
	return phone.price
}

func (phone PhoneDate) getcolor() string {
	var r string
	switch phone.color {
	case RED:
		r = "Red"
	case WHITE:
		r = "White"
	case YELLOW:
		r = "Yellow"
	default:
		r = ""
	}
	return r
}

// phone是接收器，如果使用phone PhonDate的话，是赋值语句是不会生效的
// 可以认为是传参方式，但是加上指针之后，就是传指针方式，可以保留修改效果
func (phone *PhoneDate) setname(name string) {
	phone.name = name
}

func (phone *PhoneDate) setprice(price int) {
	phone.price = price
}

func (phone *PhoneDate) setcolor(color COLOR) {
	phone.color = color
}

func (phone PhoneDate) print() {
	fmt.Printf("Name = %s, Price = %d, Color = %s\n", phone.getname(), phone.getprice(), phone.getcolor())
}

func main() {
	phone := new(PhoneDate)
	phone.setname("Huawei")
	phone.setprice(3999)
	phone.setcolor(RED)
	phone.print()
}
