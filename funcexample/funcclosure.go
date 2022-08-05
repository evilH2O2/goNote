package funcexample

import (
	"fmt"
)

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

func closure2(host string) func(string) string {
	return func(port string) string {
		s := fmt.Sprintf("%s: %s\n", host, port)
		return s
	}
}

// 返回多个函数
func closure3(rhost string) (func(int), func(string) string) {
	rh := func(port string) string {
		s := fmt.Sprintf("%s: %s\n", rhost, port)
		return s
	}

	ids := func(id int) {
		fmt.Printf("id: %d - \n", id)
	}

	return ids, rh
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

	f := closure2("127.0.0.1")
	fmt.Printf("f: %v\n", f("8080"))
	fmt.Printf("f: %v\n", f("4444"))

	f2, f3 := closure3("127.0.0.1")
	f2(1)
	// id: 1 会提前输出
	fmt.Printf("closure3: %v\n", f3("8080"))
}
