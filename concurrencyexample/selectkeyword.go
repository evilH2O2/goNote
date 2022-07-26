package concurrencyexample

import (
	"fmt"
	"time"
)

// select 类似于 swich
// select 监听 case 语句中的 channel 读写操作
// 当 case 中的 channel 读写处于非阻塞状态时，会触发相应的动作

// 如果有多个 case 都可以执行，select会随机公平的选择一个执行，其他不会执行
// 如果没有可运行的 case，且有 default，那会执行 default
// 如果 default 也不存在，那 select 会阻塞，直到某的 case 可通信运行
// default 总是可运行的

var lLocalchannel = make(chan string)

func sendChannelData() {
	defer close(lLocalchannel)

	fmt.Println("send data...")
	time.Sleep(time.Millisecond * 200)

	lLocalchannel <- "hello world"
}

func InvokeSelectChannel() {
	go sendChannelData()

	for {
		select {
		case str := <-lLocalchannel:
			fmt.Println("str -> ", str)
		default:
			// 主协程如果执行太快，会先将 default 执行
			fmt.Println("done")
		}

		time.Sleep(time.Second)
	}

}
