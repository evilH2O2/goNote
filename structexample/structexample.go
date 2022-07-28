package structexample

import "fmt"

// 定义结构体
type VictimData struct {
	Hostname string
	Ip       string
	OpenPort []int
	IsAdmin  bool
}

func SetVictim() {
	var rports []int = []int{22, 53, 80, 3389}

	// 实例化结构体
	var victimA VictimData = VictimData{
		Hostname: "dcorp-dc",
		Ip:       "127.0.0.1",
		OpenPort: rports,
		IsAdmin:  true,
	}

	fmt.Println("victim ip => ", victimA.Ip)

	// 结构体指针
	// new() 返回类型指针
	var victimB *VictimData = new(VictimData)
	victimB.IsAdmin = true

	// 或者用取值符：&
	var victimC *VictimData = &VictimData{
		Hostname: "dcorp-sql",
	}

	// 访问
	// */& 的优先级低于 . ，所以要加括号
	fmt.Println("victimC hostname =>", (*victimC).Hostname)
	// 简写。这里还是引用
	fmt.Println("victimC hostname =>", victimC.Hostname)

	// 调用含有继承的结构体
	var attackA AttackData = AttackData{
		// 直接使用
		VictimData: VictimData{
			Hostname: "kali",
			Ip:       "10.10.10.11",
		},
		IsLocal: true,
	}
	fmt.Printf("attackA => %v\n", attackA)

	var admin AdminData = AdminData{
		// 引用
		VictimData: &victimA,
	}
	fmt.Printf("admin => %v\n", admin.OpenPort)

	// 调用含有继承的结构体指针
	var admin2 *AdminData = &AdminData{
		VictimData: &VictimData{
			Hostname: "admin-pc",
		},
	}

	fmt.Println("admin2 hostname => ", (*admin2).Hostname)
	// 简写。相当于 admin2.VictimData.Hostname
	fmt.Println("admin2 hostname => ", admin2.Hostname)

	anonymousStruct()
}

// 继承
type AttackData struct {
	VictimData
	IsLocal    bool
	IsAnyLogin bool
}

type AdminData struct {
	// 引用类型也可以
	*VictimData
}

// 匿名结构体
func anonymousStruct() {
	var var1 struct {
		id   int
		name string
	}

	var1.id = 10
	var1.name = "A"

	fmt.Printf("var1: %v\n", var1)
}
