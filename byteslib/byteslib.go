package byteslib

import (
	"bytes"
	"fmt"
)

// bytes 包提供对字节切片进行读写操作的一系列函数
func bytesContainsFuncs() {
	s1 := "Hello"
	b1 := []byte(s1)

	s2 := "he"
	b2 := []byte(s2)

	// 是否包含指定字节
	// 字符串也有相同的方法
	b := bytes.Contains(b1, b2)
	fmt.Printf("b: %v\n", b)
}

// bytes.Replace()
// bytes.Join()
// bytes.Runes()

// 只需要看看官方文档就好了

func InvokeBytesExample() {
	bytesContainsFuncs()
}
