package loglib

import (
	"fmt"
	"log"
	"os"
)

func logPrint() {
	log.Print("log...")
	log.Printf("log : %d", 1)
}

func logPanic() {
	log.Println("process run...")
	defer log.Println("done...")

	log.Panicln("Fatal error...")

	log.Println("log run...")
}

func logFatal() {
	log.Println("process run...")
	defer log.Println("done...")

	log.Fatal("Fatal error...")

	log.Println("log run...")
}

func setLogPrint() {
	i := log.Flags()
	// 打印 log 配置
	fmt.Printf("i: %v\n", i)

	// 设置显示格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 设置前缀
	log.SetPrefix("[*] Process Log: ")

	log.Println("log setting...")

	// 日志输出位置
	file := "/tmp/a.log"
	f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	defer f.Close()

	log.SetOutput(f)

	log.Println("start log...")
}

func customLogSytle() {
	// var successLogFile io.Writer = os.Stdout
	// var errorLogFile io.Writer = os.Stderr

	successLog := log.New(os.Stdout, "[*] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "[-] ", log.Ldate|log.Ltime|log.Lshortfile)

	successLog.Println("run success")
	errorLog.Panicln("fatal...")
}

func InvokeLogExample() {
	// logPrint()
	// logPanic()
	// logFatal()
	// setLogPrint()
	customLogSytle()
}
