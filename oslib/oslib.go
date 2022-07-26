package oslib

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	// 仅创建文件
	f, err := os.Create("test.md")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func createDirs() {
	// 仅创建目录
	// p.txt 也是目录
	// err := os.Mkdir("test", os.ModePerm) // 单层目录
	err := os.MkdirAll("/tmp/code/test/a.txt", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除文件/目录
func deleteDirs() {
	// 可以删除目录或文件
	// 但目录只有为空才能删除
	// err := os.Remove("/tmp/code/test")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// 可以删除目录或文件
	// 即使目录不为空
	err2 := os.RemoveAll("/tmp/code/test")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 获得当前工作目录
func getPwds() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

// 改变当前工作目录
func chdirs() {
	err := os.Chdir("/tmp/code")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	getPwds()
}

// 获取用于临时文件的默认目录
func getTemps() {
	s := os.TempDir()
	fmt.Printf("temp: %v\n", s)
}

// 重命名文件
func renameTest() {
	// 重命名目录
	// err := os.Rename("/tmp/code/test", "/tmp/code/gotest")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// 重命名文件
	err2 := os.Rename("/tmp/code/gotest/a.txt", "/tmp/code/gotest/b.txt")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 文件读写
func fileRW() {
	err := os.WriteFile("/tmp/code/gotest/b.txt", []byte("hello"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	b, err2 := os.ReadFile("/tmp/code/gotest/b.txt")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("file context: %v\n", string(b[:]))
		// fmt.Printf("file context: %v\n", string(b))
	}
}

func InvokeOSLibExample() {
	// createFile()
	// createDirs()
	// deleteDirs()
	// getPwds()
	// chdirs()
	// getTemps()
	// renameTest()
	// fileRW()
}
