package funcexample

import "fmt"

// 闭包

// 函数返回一个函数
// 函数内的函数可以访问函数内的局部变量
func test() func(int) int {
	p := 0
	return func(i int) int {
		p += i
		return p
	}
}

func InvokeFunClosure() {
	// 将返回的函数赋给变量
	var1 := test()
	fmt.Printf("var1(10): %v\n", var1(10))
	// 第一次执行，p 变量被赋值为 10
	// 第二次执行的时候，因为返回的函数与 p 保持关联，可访问 p
	// 所以 p = 20

	// 此时 var1 就是个闭包，只要在 var1 的生命周期内，p 将一直有效
	fmt.Printf("var1(10): %v\n", var1(10))

	var2 := test()
	var3 := var2(10)
	fmt.Printf("var3: %v\n", var3)
	fmt.Printf("var2(10): %v\n", var2(10))
	fmt.Printf("var1(10): %v\n", var1(10))

	var4 := test()

	// var3 将从这里重新开始
	// var2 的作用域已经退出
	var3 = var4(100)
	fmt.Printf("var3: %v\n", var3)
}
