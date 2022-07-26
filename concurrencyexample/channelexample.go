package concurrencyexample

// 两个协程之间是无法直接通信的

// 通道需要指明数据类型
// 给定时间只有一个协程可以访问数据项，因此按照设计不会发生数据竞争

// 两种通道类型：无缓存通道和缓存通道
// 无缓存通道用于执行协程之间的 同步 通信
// 缓存通道用于执行协程之间的 异步 通信

// make() 创建，chan 为关键字加数据类型
// make(chan int)      // 无缓存
// make(chan int 10)   // 缓存

// 将值发送到通道用 <- 运算符
// var1 := make(chan int 10)
// var1 <- 124

// 无论发送还是接收，箭头都指向左边
// 变量在左边是 发送数据
// 变量在右边是 接收数据
// num := <-var1

// 发送/接收操作在完全完成之前是一直被阻塞的

// 通道用完后需要关闭
// defer close(var1)

import (
	"fmt"
	"math/rand"
	"time"
)

var chan1 = make(chan int)

func generateRandNumber() {
	// 种子一直变动。Intn 才会产生随机数
	rand.Seed(time.Now().UnixMicro())
	number := rand.Intn(10)
	fmt.Println("send number => ", number)
	time.Sleep(time.Second * 2)

	// 将值放入通道
	chan1 <- number
}

func GetChanValue() {
	defer close(chan1)

	// 启动一个协程
	go generateRandNumber()
	fmt.Println("wait...")

	// 等待接收，发送与接收完成之前会一直阻塞
	number := <-chan1
	fmt.Println("get number => ", number)
}
