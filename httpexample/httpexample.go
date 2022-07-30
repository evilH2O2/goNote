package httpexample

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"sync"
)

var mainWg sync.WaitGroup

func cmdHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/shell")
	io.WriteString(w, "you ip: "+r.RemoteAddr+"\n")

	// 获取 URL 参数
	s := r.URL.Query().Get("cmd")
	io.WriteString(w, "cmd = "+s)
}

func httpServer() {
	defer mainWg.Done()

	// 多个 request handle
	sm := http.NewServeMux()

	sm.HandleFunc("/shell", cmdHandle)

	sm.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// 静态文件访问
	// 路径要么是绝对路径，要么是相对运行程序的路径(main.go)
	h := http.FileServer(http.Dir("httpexample/static"))
	sm.Handle("/file/", http.StripPrefix("/file/", h))
	// sm.Handle("/", http.FileServer(http.Dir("./statics")))
	// sm.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", sm))
}

func httpClient() {
	defer mainWg.Done()
	count := 0
	for {
		r, err := http.Get("http://127.0.0.1:8080/shell?cmd=whoami")
		if err != nil {
			log.Printf("http get err: %v\n", err)
			count++
			log.Println("re connect: ", count)
			continue
		} else {
			buf := make([]byte, r.ContentLength)
			r.Body.Read(buf)
			fmt.Printf("buf: %v\n", string(buf))
			break
		}
	}
}

func InvokeHttpExample() {
	mainWg.Add(2)

	go httpServer()

	httpClient()

	mainWg.Wait()
}
