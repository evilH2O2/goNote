package flagexample

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var help *bool
var arch string
var goos string

func usage() {
	fmt.Printf("\n[!] Flag -os [windows/linux/mac] -a [x86/x64] \n\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func InvokeFlag() {
	help = flag.Bool("h", false, "help")
	// flag.BoolVar(&help, "a", false, "help")
	flag.StringVar(&goos, "os", runtime.GOOS, "OS")
	flag.StringVar(&arch, "a", runtime.GOARCH, "Arch")

	// 帮助信息
	flag.Usage = usage

	// 解析 flag
	flag.Parse()

	if *help {
		usage()
	}

	fmt.Println("OS: ", goos)
	fmt.Println("Arch: ", arch)
}
