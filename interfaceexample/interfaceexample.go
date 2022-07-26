package interfaceexample

import "fmt"

// 方法定义的集合。注意，是方法
// 接口时值类型，保存的是：值 + 原始类型
type Master interface{}

type blackDog struct {
	Name string
}

type whiteDog struct {
	Name string
}

// 没有字段，引用或拷贝类型都无所谓
func (black *blackDog) run() {
	fmt.Printf("%s is run...\n", black.Name)
}

// 有字段，并使用了。建议使用引用类型，方便修改字段，而不是值拷贝
func (white *whiteDog) run() {
	fmt.Printf("%s is run...\n", (*white).Name) // 简写 while.Name
}

// func callName()

func InvokeStrut() {
	var1 := blackDog{
		Name: "black",
	}
	var1.run()

	var2 := whiteDog{
		Name: "white",
	}
	var2.run()
}

// 实现一个场景，主人家有两只狗，一只黑色一只白色
// 现在，会叫的狗儿有饭吃，想要吃饭就得会叫
// 要求：写一个可以接收两只狗的函数，黑狗调用黑狗叫，白狗调用白狗叫
// 这在 python 中很容易实现，定义一个接收函数的参数，进行判断输出即可

// def say(dog):
// 	# 判断传入的是白狗还是黑狗
// 	if dog == "black":
// 		color = "black"
// 	print(f"{color} WwWw")
// 但，go 需要判断数据类型 python 不需要
// 所以不同结构，无法调用同一个方法，或者说，这个方法没法定义输入的数据结构:

// func say(dog 数据结构) {
// 	fmt.Println("www")
// }

// 所以出现了接口，接口不管你输入了什么数据类型，只管你实现了那些方法
// 至于函数定义没有具体实现
// 定义一个类型，一个抽象的类型，只要实现了 say() 这个方法的类型都可以称为 eat 类型
type eat interface {
	run()
	jump()
}

// 不在乎输入的数据类型
func youcanrun(dog eat) {
	dog.run()
}

func (black *blackDog) jump() {
	fmt.Println("black can jump")
}

func (white *whiteDog) jump() {
	fmt.Println("white can jump")
}

func Dogcanrun() {
	var1 := blackDog{
		Name: "black",
	}
	// 注意，引用类型
	youcanrun(&var1)
	var1.jump()

	var2 := whiteDog{
		Name: "white",
	}
	youcanrun(&var2)
	var2.jump()
}
