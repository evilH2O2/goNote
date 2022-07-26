package funcexample

import "fmt"

func newFunc(rhost string, rport int) (string, int) {
	return rhost, rport
}

// 给返回值命名
func newFunc2(rhost string, rport int) (host string, port int) {
	// 无需 := 声明，直接赋值
	host = rhost
	port = rport

	// 无需返回名，上面已经给出了
	return
}

func anonymousFunc() {
	// 像声明变量一样声明函数
	var setHost = func(lhost string) {
		fmt.Println("[*] host => ", lhost)
	}

	setHost("127.0.0.1")

	// 匿名函数
	// 上面的其实算还有名字。或者这样，不加函数名
	// 没有参数，末尾也要加一个 ()
	var1 := func(rhost string) bool {
		fmt.Println("[*] host => ", rhost)
		// if rhost != "" {
		// 	return true
		// }
		// return false
		return rhost != ""
	}("0.0.0.0")

	fmt.Println("[*] rhost => ", var1)
}

// 返回函数(返回一个匿名函数，初始化时就像新建匿名函数一样)
func returnFunc() func(int) int {
	count := 0

	return func(arg int) int {
		fmt.Println("arg => ", arg)
		count++
		fmt.Println("count => ", count)
		return count
	}
}

// defer 延迟函数调用。会在函数返回之前进行调用
// 如果有多个defer表达式，调用顺序类似于栈，越后面的defer表达式越先被调用
// 先给返回值赋值，然后调用 defer
// https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
func deferExample() {
	defer fmt.Println("hello")
	fmt.Println("world")
	defer fmt.Println("!!")
}

func InvokeNewFunc() {
	rh, rp := newFunc("127.0.0.1", 80)

	fmt.Println("rhost = ", rh)
	fmt.Println("rhost = ", rp)

	host, port := newFunc2("127.0.0.1", 80)
	fmt.Println("host = ", host)
	fmt.Println("port = ", port)

	anonymousFunc()

	// 就像新建匿名函数，只不过这个匿名函数提前被定义了，并作为一个返回值来使用
	newf := returnFunc()
	newf(12)

	deferExample()
}
