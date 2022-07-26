package concurrencyexample

import (
	"fmt"
	"sync"
)

// 协程间同步

// 创建一个 WaitGtoup 类型的变量
var wg sync.WaitGroup

func echoNumber(i int) {
	// Done() == Add(-1)
	defer wg.Done()
	fmt.Println("number => ", i)
}

func InvokeWaitGroup() {
	for i := 0; i < 5; i++ {
		go echoNumber(i)

		// 每创建一个协程，add 加一
		wg.Add(1)
	}

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("Done")
}
