package interfaceexample

import "fmt"

type Info interface {
	echoInfo()
}

type pc struct {
	name string
}

type pc2 struct {
	name string
}

// 多个结构体实现同一个接口，多态

// 指针接收者，改变字段值更方便
func (pc1 *pc) echoInfo() {
	fmt.Printf("pac1 Hostname => %s, pc1 addr => %p\n", pc1.name, pc1)
	pc1.name = "admin-pc"
}

// 值接收者
func (pctwo pc2) echoInfo() {
	fmt.Printf("pctwo Hostname => %s, pctwo addr => %p\n", pctwo.name, &pctwo)
	pctwo.name = "admin-pc"
}

func InvokePointerInterface() {
	var2 := &pc{
		name: "dcorp-pc",
	}
	fmt.Printf("var2 addr => %p\n", var2)
	var2.echoInfo()
	fmt.Printf("pac1 => %v\n", var2)

	var var3 *pc = new(pc)
	var3.name = "pss-dc"
	var3.echoInfo()

	var var4 pc2 = pc2{
		name: "pc2",
	}
	fmt.Printf("var4 addr => %p\n", &var4)
	var4.echoInfo()
	fmt.Printf("pac4 => %v\n", var4)
}
