package concurrencyexample

import (
	"fmt"
	"time"
)

// timer只执行一次
// ticker 可以周期执行

// 用在 C2 的通信上好像不错
// 设置每 2s 读取/发送数据到通道
// select 判断通道内是否存在数据
// 客户端读取通道数据

// 配合管道通信 + select 定时发送数据

func InvokeTickerExample() {
	// new ticker
	ticker1 := time.NewTicker(time.Second * 1)

	// 新的协程运行 ticker
	go func() {
		count := 0
		for s := range ticker1.C {
			// for 直接使用 _
			// 提示 simplify range expression
			_ = s
			fmt.Println("ticker run...")
			count++

			// 3 个周期结束
			if count == 3 {
				ticker1.Stop()
				break
			}
		}
	}()

	// 主协程执行自己的工作
	fmt.Println("main run...")
	for {
		fmt.Println("main work...")
		time.Sleep(time.Second * 3)
	}
}
