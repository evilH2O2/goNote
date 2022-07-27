package oslib

import (
	"fmt"
	"os"
	"time"
)

// 进程相关
func InvokeOSProcess() {
	fmt.Printf("pid: %v\n", os.Getpid())
	fmt.Printf("ppid: %v\n", os.Getppid())

	// cretae new process
	// 设置新的进程属性
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		// 环境变量
		Env: os.Environ(),
	}

	p, err := os.StartProcess("/System/Applications/Calculator.app/Contents/MacOS/Calculator", nil, &attr)
	if err != nil {
		fmt.Printf("StartProcess err: %v\n", err)
	} else {
		fmt.Printf("New Process Pid: %v\n", p.Pid)
	}

	// 通过 PID 查找进程
	p2, err2 := os.FindProcess(p.Pid)
	if err2 != nil {
		fmt.Printf("FindProcess err: %v\n", err2)
	} else {
		time.AfterFunc(time.Second*8, func() {
			// 发送退出信号
			p2.Signal(os.Kill)
		})

		// 等待进程退出并返回状态
		ps, err3 := p.Wait()
		if err3 != nil {
			fmt.Printf("Process Wait err: %v\n", err3)
		} else {
			fmt.Printf("ps.ExitCode(): %v\n", ps.ExitCode())
			fmt.Printf("ps.String(): %v\n", ps.String())
		}
	}
}
