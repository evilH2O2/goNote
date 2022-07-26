package processcontrol

import "fmt"

func FunIfElse() {
	var count int = 100
	if count > 100 {
		fmt.Println("count > 100")
	} else if count == 100 {
		fmt.Println("count == 100")
	}
}

func FunSwichCase() {
	// case 结尾自动 break，如果要继续匹配下一个可以加入 fallthrough
	var var1 int = 100
	var var2 int

	switch {
	case var1 < 10:
		var2 = 0
		fmt.Printf("count < 10")
	case var1 > 10:
		var2 = 1
		fmt.Println("count > 10")
		// 继续匹配下一项
		fallthrough
	case var1 == 100:
		var2 = 9
		fmt.Println("count == 100")
		// 自动 break
	}

	switch var2 {
	case 1:
		fmt.Printf("var2 == 1")
	case 2:
		fmt.Printf("var2 == 2")
	case 9:
		fmt.Printf("var2 == 9")
	default:
		fmt.Printf("var2 = ?")
	}
}

func FunFor() {
	var1 := 10

	// 无限循环
	for {
		fmt.Printf("var1 => %d\n", var1)
		var1++
		if var1 == 15 {
			break
		}
	}

	// 条件循环
	for var1 < 20 {
		fmt.Printf("var2 => %d\n", var1)
		var1++
	}

	// 标准 for 循环
	for v := 18; v < var1; v++ {
		fmt.Printf("v => %d\n", v)
	}
}
