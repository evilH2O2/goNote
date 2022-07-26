package concurrencyexample

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
// 除了使用通道进行同步外，还可以使用 Mutex 互斥锁进行同步

// 争夺一个共享变量
var i int = 100
var mwg sync.WaitGroup

// 新建一个互斥锁
var lLock sync.Mutex

func add() {
	defer mwg.Done()
	// 运行前加锁
	// 一个协程运行前加锁，另一个协程运行时是无法进入代码执行(被 lock 了)的
	// 只有等解锁后，另一个协程才会执行然后加锁

	// 注意，加锁前的代码时可以运行的
	fmt.Println("add wait...")

	lLock.Lock()
	i += 1
	fmt.Println("i++ => ", i)
	time.Sleep(time.Millisecond * 200)
	// 运行结束解锁
	lLock.Unlock()
}

func sub() {
	defer mwg.Done()
	// 加锁
	fmt.Println("sub wait...")
	lLock.Lock()
	i -= 1
	fmt.Println("i-- => ", i)
	time.Sleep(time.Millisecond * 50)
	// 解锁
	lLock.Unlock()
}

func InvokeMutex() {
	for i := 0; i < 100; i++ {
		mwg.Add(1)
		go add()
		mwg.Add(1)
		go sub()
	}

	mwg.Wait()
	fmt.Printf("done i: %v\n", i)
}
