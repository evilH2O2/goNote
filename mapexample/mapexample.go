package mapexample

import "fmt"

// map / 映射
// 无序键值对。不分配内存时，默认值为 nil。容量自动增长
// make 分配内存空间
func MapExample() {
	// 键为基本数据类型，值可以为任何数据类型
	var var1 map[string]int

	fmt.Println("var1 does nil? => ", var1 == nil)

	var1 = make(map[string]int)
	var1["init"] = 0
	var1["port"] = 1
	fmt.Println("var1 => ", var1)

	// 简写声明赋值
	var2 := map[string]string{
		"host": "127.0.0.1",
		"port": "8080",
	}
	fmt.Println("var2 => ", var2)

	// 取值
	// 不加判断，键不存在会返回空，所以可以加一个判断
	var1Value := var2["host"]
	fmt.Println("var2[\"host\"] => ", var1Value)

	// 加判断
	var1ValueNull, isExist := var2["hosts"]
	if isExist {
		fmt.Println("var2[\"host\"] => ", var1ValueNull)
	} else {
		fmt.Println("var2[\"hosts\"] is not exist")
	}

	// 删除键值对
	delete(var2, "host")
	fmt.Println("var2 => ", var2)

	// 枚举
	for i, v := range var1 {
		fmt.Printf("key => %s, value => %d\n", i, v)
	}
}
