package errorslib

import (
	"fmt"
	"log"
)

// panic 捕获
// 不会立即退出，最后调用 defer
// 详解：https://www.v2ex.com/t/725736
func InvokeErrorsRecover() {
	// panic 错误会执行 defer
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()

	log.Panicln("panic...")
	fmt.Println("End...")
}
