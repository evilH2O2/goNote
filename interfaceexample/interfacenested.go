package interfaceexample

import "fmt"

// 接口嵌套

// 接口通过嵌套，创建新的接口
type Run interface {
	run()
}

type Jump interface {
	jump()
}

// 接口嵌套
type Dogs interface {
	Run
	Jump
}

// 定义结构体
type dog struct {
	name string
}

// 实现嵌套接口
func (d *dog) run() {
	fmt.Printf("%s can run\n", (*d).name)
}

func (d *dog) jump() {
	fmt.Printf("%s can jump\n", (*d).name)
}

func InvokeInterfaceNested() {
	// dog := new(dog)
	var dog *dog = new(dog)
	dog.name = "alis"
	dog.run()
	dog.jump()
}
