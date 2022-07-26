package concurrencyexample

import (
	"fmt"
	"time"
)

// 可以实现一些定时操作，内部通过 channel 实现

// 创建方式:
// time.NewTimer()
// time.After()
func InvokeTimerExample() {
	// 创建一个定时器
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("Now Time: ", time.Now())

	// C 其实是个 channel 类型的
	// 阻塞指定时间后读取
	t1 := <-timer1.C
	fmt.Println("t1 => ", t1)

	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C

	// 单纯等待时间，可以使用 time.Sleep()
	time.Sleep(time.Second * 2)

	// 指定时间之后执行
	<-time.After(time.Second * 2)

	timer3 := time.NewTimer(time.Second * 3)
	go echoHello(timer3)

	// 将不再执行定时任务
	sstatus := timer3.Stop()
	if sstatus {
		fmt.Println("timer stop...")
	}

	time.Sleep(time.Second * 2)
	fmt.Println("main stop...")
}

func echoHello(t *time.Timer) {
	fmt.Println("wait 3s...")
	// 定时器停止后，下面的代码将不再执行
	<-t.C
	fmt.Println("hello")
}
