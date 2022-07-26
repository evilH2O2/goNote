package methodexample

import "fmt"

// 定义结构体
type VictimData struct {
	Hostname string
	Ip       string
	OpenPort []int
	IsAdmin  bool
}

// 方法。跟定义函数差不多
// 但方法必须要有一个接收者，左边的那一部分
func (victim VictimData) enumVictimData() {
	fmt.Println("Hostname => ", victim.Hostname)
	fmt.Println("IP => ", victim.Ip)
	fmt.Println("Open Ports => ", victim.OpenPort)
	fmt.Println("IsAdmin => ", victim.IsAdmin)
}

// 可以接收指针
func (victim *VictimData) setIp() {
	// 可以简写。victim.Ip
	(*victim).Ip = "10.10.10.12"
}

// 调用方法
func VictimConfig() {
	var rports []int = []int{22, 53, 80, 3389}

	// 实例化结构体
	var victim VictimData = VictimData{
		Hostname: "dcorp-dc",
		Ip:       "127.0.0.1",
		OpenPort: rports,
		IsAdmin:  true,
	}

	victim.setIp()
	victim.enumVictimData()
}
