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
}
