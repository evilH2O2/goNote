package concurrencyexample

import (
	"fmt"
	"sync/atomic"
)

// 除了加锁 sync.Mutex 外
// 其他的防止公共资源争夺导致数据错乱的方法

// 通过 atomic 保护数据
// 原子变量
// 善用 atomic 能够避免程序中出现大量锁操作

// 注意数据类型
var num int32 = 100

func adds() {
	// num++
	// AddInt32 在原址的基础上增加
	atomic.AddInt32(&num, 1)
}

func subs() {
	// num--
	atomic.AddInt32(&num, -1)
}

func InvokeAtomicExample() {
	go adds()
	go subs()

	fmt.Println("num => ", num)
}
