package concurrencyexample

import (
	"fmt"
	"runtime"
	"time"
)

func echoNum(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("number => ", i)
		time.Sleep(time.Millisecond * 50)
		if i == 3 {
			runtime.Goexit()
		}
	}
}

func InvokeRuntime() {
	go echoNum(5)
	for i := 0; i < 4; i++ {
		// 让给其他协程执行
		runtime.Gosched()
		time.Sleep(time.Millisecond * 100)
		fmt.Println("master run => ", i)
	}

	fmt.Println("main run...")
	// fmt.Println(runtime.NumCPU())
}
