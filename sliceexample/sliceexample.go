package sliceexample

import "fmt"

// 切片是对数组的引用，默认值为 nil
// 因为是引用类型，改了 切片 的值相当于改了 数组 的值
func SliceExample() {
	var1 := [4]int{1, 2, 3, 4}
	var var2 []int = var1[:]

	var2[0] = 2
	fmt.Println(var1)

	// 切片引用切片的一部分(引用的都是同一个底层数组)
	var3 := var2[:]
	var3[1] = 8

	fmt.Println("var1 => ", var1)
	fmt.Println("var2 => ", var2)
	fmt.Println("var3 => ", var3)
	// fmt.Printf("var1 addr => %p\nvar2 addr => %p\nvar3 addr => %p\n", &var1, &var2, &var3)

	var var4 []int
	fmt.Println("var4 does nil => ", var4 == nil)

	// nil 是没有地址空间的，现在 var4 无法被赋值，需要分配内存空间
	// 可以放 3 个数据，最长为 5 个，默认容量与长度相等
	var4 = make([]int, 3, 5)
	fmt.Printf("var4 len => %d, cap => %d\n", len(var4), cap(var4))

	// 追加内容。底层创建新数组，不再引用原数组
	var2 = append(var2, 5, 6, 7, 8)
	fmt.Printf("var2 addr => %p, var2 => %v\n", &var2, var2)

	var2[0] = 9
	fmt.Printf("var1 => addr => %p, %v\n", &var1, var1)

	// 底层元数组已经不在变化
	fmt.Printf("var2 => addr => %p, %v\n", &var2, var2)

	// 追加切片
	var5 := append(var2, var3...)
	fmt.Println("var5 => ", var5)

	// 复制数组。数据不够，用 0 填满(其实就是未将 0 覆盖完)
	var6 := make([]int, 9)
	copy(var6, var2)
	fmt.Println("var6 copy => ", var6)

	// string 与 []byte 互转
	var7 := "hello world"
	fmt.Printf("var7 []byte => %v\nvar string => %s\n", []byte(var7), []byte(var7))

	var8 := []int{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	for _, v := range var8 {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	sliceFunc("hello", "how are you?")
}

// 函数接收多个参数
func sliceFunc(arg ...string) {
	for _, v := range arg {
		fmt.Println("v => ", v)
	}
}
