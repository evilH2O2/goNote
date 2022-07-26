package concurrencyexample

import (
	"fmt"
	"sync/atomic"
)

// AddInt32/AddInt64 在原值的基础上增加
func addAdnSub() {
	var i int32 = 10
	atomic.AddInt32(&i, 1)
	fmt.Println("i AddInt32 => ", i)
}

func readAndWrite() {
	// read
	var i int32 = 10
	atomic.LoadInt32(&i)
	fmt.Printf("i: %v\n", i)

	// write
	atomic.StoreInt32(&i, 200)
	fmt.Printf("i: %v\n", i)
}

func cas() {
	// AddInt/LoadInt/StoreIn 的底层
	i := int32(100)

	// 比较并交换
	// 返回 bool 类型，交换是否成功
	// 比较指定位置的值是否与旧值相同（是否有别的协程/线程将值修改了）
	// 如果被修改了，对比不成功，则放弃交换，返回 false
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b: %v, i: %v\n", b, i)
}

func InvokeAtomicUse() {
	addAdnSub()
	readAndWrite()
	cas()
}
