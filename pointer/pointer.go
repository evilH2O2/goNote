package pointer

import (
	"fmt"
)

func numberAdd(num int) {
	// 值拷贝(值类型)。变量开辟新地址，将传入的值复制到新地址
	num++
	// & 取址符
	// * 取值符
	// *数据类型 指向的类型
	fmt.Printf("num => %d, addr => %p\n", num, &num)
}

func numberAddPointer(num2 *int) {
	// 值引用(引用类型)。变量开辟新地址，将变量地址放入新地址，直接改变这个地址内的值

	// * 取值符，从地址中取出值进行操作
	*num2++
	fmt.Printf("num2 => %d, addr => %p\n", *num2, &num2)
}

func arrayPointer() {
	// golang 指向数组的指针与 C 不同
	// golang 数组中的每个值都是指针类型
	// C 赋值时是取数组首地址，而 golang需要每个都赋值

	var1 := [3]int{1, 2, 3}

	ptr1 := [3]*int{}

	// 不能直接 ptr1 = &var1
	for i := 0; i < len(var1); i++ {
		ptr1[i] = &var1[i]
	}

	fmt.Printf("ptr1: %v\n", ptr1)
	fmt.Printf("ptr1: %T\n", ptr1[0])
	fmt.Printf("ptr1: %v\n", *ptr1[0])
}

func PointerVar1() {
	var number int = 0
	var numberAddr *int = &number
	// 空指针，= nil ，new() 系统分配一块内存，值为 0
	var nuiPionter *int = new(int)

	numberAdd(number)
	numberAddPointer(numberAddr)
	// numberAddPointer(&number) 可以直接传入地址

	fmt.Printf("number => %d, addr => %p\n", number, &number)
	fmt.Printf("now number => %d, addr => %p\n", number, &number)
	fmt.Printf("nuiPionter => %v, value => %d, addr => %p\n", nuiPionter, *nuiPionter, &nuiPionter)

	arrayPointer()
}

func echoNumber(num *int) {

	// num 中存的是一个地址，要是取值的话，需要取值符 *num
	// 指针变量在使用的时候，非特殊情况都是需要取值符的 *
	// 因为你要用它的值进行操作
	fmt.Println("number", *num)

	// 这里取的是新变量 num 的地址，而不是 num 保存的那个参数的地址
	// 所以是 &num，而不是 num
	fmt.Printf("number addr %p\n", &num)

	// 这是参数的值保存的地址，对照 ptrNum 的地址，是一样的
	fmt.Printf("num addr %p\n", num)
}

func PointerVar2() {
	var ptrNum int = 10
	var number *int = &ptrNum
	fmt.Printf("ptrNum addr: %p\n", &ptrNum)

	// 指针类型未初始化?赋值?之前是不能直接使用的
	// 大概是未初始化，这么说
	// 报错：panic: runtime error: invalid memory address or nil pointer dereference
	// 就是这样:

	// var ptrA *int
	// echoNumber(ptrA)

	// 函数的参数需要的是一个值，一个 int 指针类型的值
	// 因为 num 被声明为 *int
	// number 也被声明为 *int
	// 所以直接传递 number 就可以

	// 因为传递的是 ptrNum 的地址，所以不要加取值符 *number
	// 传递的是地址而不是值
	// 因为后续在函数内的操作是直接作用在 ptrNum 上的
	// 只是需要值的话，用值拷贝就是了，不然为什么用指针变量
	echoNumber(number)
}
