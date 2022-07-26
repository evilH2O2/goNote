package interfaceexample

import "fmt"

// 忘掉其他编程语言，这个接口就简单了
// 没有为什么，就是因为这样而已

// 一个接口可以被多个结构体实现

// 定义接口
type BaseData interface {
	setHost()
	setPort()
}

// 定义结构体
type victim struct {
	name string
}

type attack struct {
	name string
}

// 实现接口方法
// 注意接收者是那个结构体的

// victim 实现的 baseData 接口方法
// 指针接收者
func (v *victim) setHost() {
	var host string = "10.10.10.11"
	fmt.Printf("Hostname: %s, Host => %s\n", (*v).name, host)
}

func (v *victim) setPort() {
	var port int = 8080
	fmt.Printf("Hostname: %s, Port => %d\n", (*v).name, port)
}

// attack 实现的 baseData 接口方法
func (a *attack) setHost() {
	var host string = "127.0.0.1"
	fmt.Printf("Hostname: %s, Host => %s\n", (*a).name, host)
}

func (a *attack) setPort() {
	var port int = 8080
	fmt.Printf("Hostname: %s, Port => %d\n", (*a).name, port)
}

func InvkeInterfaceMethod() {
	v := &victim{
		name: "dcorp-dc",
	}

	// 调用接口方法
	v.setHost()
	v.setPort()

	a := &attack{
		name: "kali",
	}

	// 实现接口必须实现接口中的所有方法
	a.setHost()
	a.setPort()
}
