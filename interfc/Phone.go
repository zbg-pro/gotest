package interfc

import (
	"fmt"
	"time"
)

type Phone interface {
	Call()
	Receive() string
	Count() int
	Calc(x int, y int) (int, string)
	SetName(x string)
}

type NokiaPhone struct {
	company string
	time    time.Time
	Name    string
}

func (nokiaPhone *NokiaPhone) SetName(x string) {
	//TODO implement me
	nokiaPhone.Name = x
}

func (nokiaPhone NokiaPhone) Count() int {
	//TODO implement me
	return 2
}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("Hello I am nokia!")
}
func (nokiaPhone NokiaPhone) Receive() string {
	fmt.Println("revive: nokia phone response")
	return "nokia phone response" + nokiaPhone.company
}

func (nokiaPhone NokiaPhone) Calc(x int, y int) (int, string) {
	rs := x + y
	return rs, " nokiaPhone result"
}

func (nokiaPhone *NokiaPhone) ChangeName(name string) {
	if name == "" {
		name = "5320"
	}
	nokiaPhone.Name = name
}

type IPhone struct {
	Version string
	Name    string
}

func (iPhone IPhone) SetName(x string) {
	//TODO implement me
	iPhone.Name = x
}

func (iPhone IPhone) Receive() string {
	return "IPhone Receive" + iPhone.Version
}

func (iPhone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}

func (iPhone IPhone) Calc(x int, y int) (int, string) {
	rs := x + y*2
	return rs, " iPhone result"
}

func (iPhone IPhone) Count() int {
	//TODO implement me
	return 1
}

func (iPhone *IPhone) ChangeName(name string) {
	if name == "" {
		name = "14plus"
	}
	iPhone.Name = name
}
