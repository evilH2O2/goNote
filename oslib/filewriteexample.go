package oslib

import (
	"fmt"
	"os"
)

// Fiel 写操作

// 写数组
func writeToByte(file string) {
	// OpenFile
	// 比 Open 要多一些功能
	// O_CREATE: 如果文件不存在，那就创建它
	// 默认覆写, os.O_APPEND 追加写入
	// 指定覆写：os.O_TRUNC
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())

		f.Write([]byte("hello pwn!"))
		f.Close()
	}

}

func readFromFile(file string) {
	buf := make([]byte, 10)

	// 获取文件句柄
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())

		// 读文件
		_, err2 := f.Read(buf)
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		}

		fmt.Printf("string(buf): %v\n", string(buf[:]))
		f.Close()
	}
}

// 直接写入字符串
func writeToString(file string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE/os.O_TRUNC, 0755)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())

		f.WriteString("hello pwn!")
		f.Close()
	}

}

// 偏移指针：f.WriteAt()

func InvokeOSFileWrite() {
	file := "/tmp/code/test/a.txt"

	writeToByte(file)
	readFromFile(file)
	writeToString(file)
	readFromFile(file)
}
