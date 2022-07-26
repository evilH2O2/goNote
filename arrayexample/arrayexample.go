package arrayexample

import "fmt"

// Array 长度固定，元素固定的一组数据

func ArrayVar1() {
	var var1 [3]int = [3]int{0, 1}
	// 通过索引赋值
	var1[2] = 3

	fmt.Println(var1[0])

	// [...] 自动推断长度
	var var2 = [...]int{123, 456, 789}
	fmt.Println(var2)

	for _, var2 := range var2 {
		fmt.Println(var2)
	}

	var twoArray [2][3]int = [2][3]int{
		{11, 12, 13},
		{14, 15, 16},
	}

	fmt.Println(twoArray)

	// v 是第二个 arrar，也就是第一层 array 的值
	// v2 是第二层的值
	for i, v := range twoArray {
		for i2, v2 := range v {
			fmt.Printf("twoArray[%d][%d] = %d\n", i, i2, v2)
		}
	}
}
