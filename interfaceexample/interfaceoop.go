package interfaceexample

import "fmt"

// 通过结构体和函数绑定来实现 OOP

// 定义一个类
type doggs struct {
	name string
	old  int
}

// 定义类方法
// 设置接收者为结构体进行绑定
func (d doggs) run() {
	fmt.Println("I can run")
}

func (d doggs) jump() {
	fmt.Println("I can jump")
}

func InvokeOOP() {
	// 实例化
	dog := doggs{
		name: "alis",
		old:  3,
	}

	fmt.Printf("%v\n", dog)
	dog.run()
	dog.jump()
}
