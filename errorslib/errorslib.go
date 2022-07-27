package errorslib

import (
	"errors"
	"fmt"
)

// 实现了操作错误的一系列函数
// error值为 nil，则表示未遇到错误

// type error interface {
// Error() string
// }

// 你可以用任何类型区实现它
// 只要添加一个 error() 方法即可

// error 不一定表示一个错误，它可以表示任何信息
// 例如 io 包中用 error 类型 io.EOF 表示数据读取结束

// error 包实现了一个最简单的 error 类型，只包含一个字符串
// 它可以记录大多情况下遇到的错误信息

// errors包只有一个 new() 函数，用于生成一个最简单的 error 对象

func checkInt(v interface{}) (interface{}, error) {
	// 接口接收的类型
	switch var1 := v.(type) {
	case int:
		err := errors.New("int in blacklist")
		return nil, err
	default:
		// fmt.Printf("var1: %T\n", var1)
		return var1, nil
	}
}

// 自定义错误
type flagQ struct {
	Flag string
	Msg  string
}

// 实现了 Error 方法相当于实现了 error 接口
func (f *flagQ) Error() string {
	return fmt.Sprintf("%s: %s", f.Flag, f.Msg)
}

// 使用这个 error
func invokeFlag(flag string, msg string) error {
	// 打印 error 信息
	return &flagQ{flag, msg}
}

// error 接口就是用来花样打印错误消息的，也只是用来打印消息的
// 其他的表达式呀语句呀都在你需要这个 error 的函数或其他什么对象里实现的

// 声明一个自定义错误信息结构
// 实现这个接口
// 使用这个接口

func InvokeErrorsExample() {
	var p int

	i, err := checkInt(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("i: %T\n", i)
	}

	flag := false
	if !flag {
		// 使用这个自定义 error
		err2 := invokeFlag("N", "flag is false")
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		}
	}
}
