package main

import "fmt"

func PrintFormatExample() {
	var var1 int = 1290
	var var2 float64 = 1290.0
	var var3 string = "how are you?"

	// %v 值( int/string 通用，但不建议)， %T 值类型
	fmt.Printf("var1 => %v => type => %T => %#v\n", var1, var1, var1)

	// 10进制整数: %d
	fmt.Printf("var1 => %d\n", var1)

	// 2进制打印
	fmt.Printf("var1 bin => %b\n", var1)

	// 输出与整数对应的字符
	fmt.Printf("var1 char => %c\n", var1)

	// 输出十六进制
	fmt.Printf("var1 hex => %x\n", var1) // 小写
	fmt.Printf("var1 hex => %X\n", var1) // 大写

	// 字符值加单引号
	fmt.Printf("var1 %%q => %q\n", var1)

	// 输出浮点值
	fmt.Printf("var1 floate => %f\n", var2)

	// 输出字符串
	fmt.Printf("var3 string => %s\n", var3)

	// 字符串以 base16 编码输出
	fmt.Printf("var3 base16 => %x\n", var3)

	// 输出指针地址
	fmt.Printf("var1 pointer => %p\n", &var1)

	// 控制整数宽度
	fmt.Printf("var1 int |%6d|\n", var1)

	// 浮点数控制精度
	fmt.Printf("var2 floate %6.2f\n", var2)

	// bool值
	fmt.Printf("true => %t\n", true)
}
