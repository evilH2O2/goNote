package structexample

import "fmt"

// 结构体作为函数参数
type data struct {
	id   int
	name string
}

func invokeStructFuncValue(p data) {
	// 值拷贝
	// 修改值不会影响外面的结构体值
	p.id = 10
	p.name = "lingling"

	fmt.Printf("p: %v\n", p)
}

func invokeStructFuncPointer(p *data) {
	// 引用传递
	// 修改值会影响外面的结构体值
	p.id = 20
	p.name = "xiaohong"

	fmt.Printf("p: %v\n", p)
	// fmt.Printf("p: %v\n", *p)
}

func InvokeStructFunc() {
	var1 := data{1, "daming"}
	invokeStructFuncValue(data{})
	invokeStructFuncValue(var1)
	invokeStructFuncPointer(&var1)

	fmt.Printf("var1: %v\n", var1)

}
