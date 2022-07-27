package oslib

import (
	"fmt"
	"io"
	"os"
)

// File 文件读操作

// 文件打开操作，获取 File 句柄
func fileOpenClose() {
	// 打开成功，返回一个 file 句柄
	var lfile string = "/tmp/code/gotest/b.txt"
	f, err := os.Open(lfile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}

	// 更多功能
	f2, err2 := os.OpenFile(lfile, os.O_CREATE|os.O_RDWR, 0755)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
		f2.Close()
	}

	// 创建临时文件，是文件而不是目录
	f3, err3 := os.CreateTemp("/tmp/code", "*")
	if err3 != nil {
		fmt.Printf("err3: %v\n", err3)
	} else {
		fmt.Printf("f3.Name(): %v\n", f3.Name())
		f3.Close()
	}
}

// 读文件
func osFileRead() {
	var lfile string = "/tmp/code/gotest/b.txt"

	// 打开一个文件句柄
	f, _ := os.Open(lfile)
	defer f.Close()

	// 新建空缓冲区，存放读取的数据
	b := make([]byte, 10)

	// 将数据读入缓冲区
	n, err := f.Read(b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("data byte size: %v\n", n)
		fmt.Printf("string(n): %v\n", string(b))
	}
	// f.Close()

	// 循环读取所有内容
	for {
		n2, err2 := f.Read(b)
		if err2 == io.EOF {
			break
		}
		fmt.Printf("read data size: %d\n", n2)
		fmt.Printf("string(b): %v\n", string(b))
	}

	// 偏移指针
	// 从第 5 字节开始读满缓冲区
	n2, _ := f.ReadAt(b, 5)
	fmt.Printf("ReadAt size: %v\n", n2)
	fmt.Printf("string(b): %v\n", string(b))

	// 读取目录，目录枚举
	// 注意深度
	de, _ := os.ReadDir("/tmp/code")
	fmt.Printf("Dirs: %v\n", de)
	for _, de2 := range de {
		fmt.Printf("de2.IsDir(): %v\n", de2.IsDir())
		fmt.Printf("de2.Name(): %v\n", de2.Name())
	}

	b = make([]byte, 5) // 重置

	// 自定义指针位置
	// 第一个参数：偏移量
	// 第二个参数：0: 开头 1: 相对位置 2: 末尾
	f.Seek(-4, 2)
	n3, _ := f.Read(b)
	fmt.Printf("n3: %v\n", n3)
	fmt.Printf("string(b): %v\n", string(b))
}

func InvokeOSFile() {
	// fileOpenClose()
	osFileRead()
}
