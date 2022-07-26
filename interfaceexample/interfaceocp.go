package interfaceexample

import (
	"fmt"
)

// OCP 原则
// 对拓展是开放的，对修改是关闭的

// 添加新功能的时候是以拓展的方式加入，而不是修改以前的代码来适应新功能
type victims interface {
	setHost(string)
}

type machine struct{}

func (m *machine) setHost(host string) {
	fmt.Printf("host => %s\n", host)
}

type control struct {
	host string
}

// victims 接口在上面已经实现了
// 以 victims 接口为参数，可以传递 machine
func (c *control) controlMachine(mhe victims) {
	// 调用了接口方法
	mhe.setHost(c.host)
}

func RunMachine() {
	ma := new(machine)
	ma.setHost("10.10.10.11")

	cl := new(control)
	cl.host = "0.0.0.0"

	// 可以传递 machine
	cl.controlMachine(ma)
}
