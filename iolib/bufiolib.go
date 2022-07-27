package iolib

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// 缓冲 I/O
// 字符串读取或者文件读取都可以

func newBufioReader() {
	// 新建 io.Reader
	r := strings.NewReader("hello")

	// 转为 bufio.Reader
	r2 := bufio.NewReader(r)
	// 读取缓冲区里的内容
	// 以 \n 为终止符
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func newBufioWriter() {
	file := "/tmp/code/test/a.txt"
	f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer f.Close()

	w := bufio.NewWriterSize(f, 5)
	i, err := w.WriteString("pwn")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("write size: %v\n", i)
	}

	// 刷新缓冲区，将缓冲区内容写入内存
	w.Flush()
}

func resetBuffer() {
	b := bytes.NewBuffer(make([]byte, 10))
	w := bufio.NewWriter(b)
	w.WriteString("hello")

	// 第二个缓冲区
	b2 := bytes.NewBuffer(make([]byte, 5))
	// 丢弃缓冲区内容
	// 将 b 输出重设为 b2
	w.Reset(b2)
	w.WriteString("world")
	w.Flush()

	fmt.Printf("b: %v\n", b)
	fmt.Printf("b2: %v\n", b2)
}

func bufferSize() {
	b := bytes.NewBuffer(make([]byte, 10))
	w := bufio.NewWriterSize(b, 10)

	w.WriteString("hello")

	// 缓冲区未使用的大小
	fmt.Printf("w.Available(): %v\n", w.Available())

	// 缓冲区已使用的大小
	fmt.Printf("w.Buffered(): %v\n", w.Buffered())

	w.Flush()
	fmt.Printf("w.Available(): %v\n", w.Available())
	fmt.Printf("w.Buffered(): %v\n", w.Buffered())

}

func InvokeBufioExample() {
	// newBufioReader()
	// newBufioWriter()
	// resetBuffer()
	bufferSize()
}
