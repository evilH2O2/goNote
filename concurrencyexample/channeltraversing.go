package concurrencyexample

import "fmt"

// channel 遍历

var localchannel = make(chan int)

func InvokeChannelTraversing() {

	go func() {
		for i := 0; i < 2; i++ {
			localchannel <- i
		}
		// 写入完成后关闭通道
		close(localchannel)
	}()

	// 读取大于写入，且通道写入后未关闭
	for v := 0; v < 3; v++ {
		var readvalue int = <-localchannel
		fmt.Println("localchannel value => ", readvalue)
	}
}
