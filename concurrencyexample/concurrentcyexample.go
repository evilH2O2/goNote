package concurrencyexample

import (
	"fmt"
	"time"
)

// 协程，比线程更轻量级

func echoMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %s => %d\n", msg, i)
		time.Sleep(time.Second * 1)
	}
}

func InvokeConcurrentcy() {
	// 启动一个协程来执行函数
	go echoMsg("A")

	// 依旧由 main 主协程执行
	echoMsg("B")
}
