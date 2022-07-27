package oslib

import (
	"fmt"
	"os"
)

func GetALLEnviron() {
	// 获取所有环境变量
	s := os.Environ()
	fmt.Printf("s: %v\n", s[0])

	// 获取某个环境变量
	s2 := os.Getenv("SHELL")
	fmt.Printf("Env SHELL: %v\n", s2)

	// 设置环境变量
	err := os.Setenv("var1", "123")
	if err != nil {
		fmt.Printf("setenv err: %v\n", err)
	}

	// 查找环境变量
	s3, b := os.LookupEnv("var1")
	if b {
		fmt.Printf("var1 found: %v\n", s3)
	}

	// 替换
	// 将右边的值替换左边的值
	// 左边的值类似占位符
	os.Setenv("rhost", "0.0.0.1")
	os.Setenv("port", "80")
	fmt.Println(os.ExpandEnv("${rhost}:${port}"))

	// 清空环境变量
	// os.Clearenv()
}
