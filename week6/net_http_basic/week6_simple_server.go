package main

import (
	"fmt"
	"log"      // 用于日志记录
	"net/http" // 导入 net/http 包
	"time"
)

// --- 1. 定义处理器函数 (Handler Functions) ---
// 处理器函数是响应 HTTP 请求的函数。
// 它必须满足 http.HandlerFunc 类型，即 func(w http.ResponseWriter, r *http.Request)。
// - http.ResponseWriter: 用于构建和发送 HTTP 响应给客户端。
// - *http.Request: 代表客户端发送的 HTTP 请求。

// helloHandler 响应 "/hello" 路径的请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path 可以获取请求的路径
	fmt.Printf("  [Server] 收到请求: %s %s\n", r.Method, r.URL.Path)

	// 设置响应头 (可选，例如 Content-Type)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// 设置响应状态码 (可选，默认为 http.StatusOK 即 200)
	// w.WriteHeader(http.StatusOK) // 如果不设置，默认就是200

	// 写入响应体
	// Fprintf 类似于 Printf，但它写入到一个 io.Writer (http.ResponseWriter 实现了 io.Writer)
	fmt.Fprintf(w, "你好，世界！Hello, World from Go HTTP Server! Requested path: %s", r.URL.Path)
}

// timeHandler 响应 "/time" 路径的请求，返回当前时间
func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("  [Server] 收到请求: %s %s\n", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
	fmt.Fprintf(w, "当前服务器时间是: %s", currentTime)
}

// headersHandler 响应 "/headers" 路径，打印请求头
func headersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("  [Server] 收到请求: %s %s\n", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "--- 请求头信息 ---")
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	fmt.Println("--- 第6周学习：net/http 标准库 (简单 HTTP 服务器) ---")

	// --- 2. 注册处理器函数 ---
	// http.HandleFunc 函数将一个处理器函数注册到指定的请求路径。
	// 当服务器收到匹配该路径的请求时，对应的处理器函数就会被调用。
	http.HandleFunc("/hello", helloHandler) // 当访问 "/hello" 时，调用 helloHandler
	http.HandleFunc("/time", timeHandler)   // 当访问 "/time" 时，调用 timeHandler
	http.HandleFunc("/headers", headersHandler)

	// 根路径 "/" 的处理器 (可以捕获所有未被其他模式匹配的请求，如果放在最后)
	// 或者作为默认欢迎页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 如果路径不是根路径，并且没有被其他处理器匹配，可以返回 404
		if r.URL.Path != "/" {
			http.NotFound(w, r) // http.NotFound 是一个便捷函数，发送 404 Not Found 响应
			fmt.Printf("  [Server] 404 Not Found: %s %s\n", r.Method, r.URL.Path)
			return
		}
		fmt.Printf("  [Server] 收到请求: %s %s (根路径)\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, "<h1>欢迎来到 Go HTTP 服务器!</h1>")
		fmt.Fprintln(w, "<p>尝试访问以下路径:</p>")
		fmt.Fprintln(w, "<ul>")
		fmt.Fprintln(w, "  <li><a href=\"/hello\">/hello</a></li>")
		fmt.Fprintln(w, "  <li><a href=\"/time\">/time</a></li>")
		fmt.Fprintln(w, "  <li><a href=\"/headers\">/headers</a></li>")
		fmt.Fprintln(w, "</ul>")
	})

	// --- 3. 启动 HTTP 服务器 ---
	// http.ListenAndServe 函数启动一个 HTTP 服务器，监听指定的 TCP 地址和端口。
	// 第一个参数是服务器地址 (例如 ":8080" 表示监听所有网络接口的 8080 端口)。
	// 第二个参数是一个 http.Handler。如果为 nil，则使用 DefaultServeMux (默认的请求路由器)，
	// 我们通过 http.HandleFunc 注册的处理器就是注册到了 DefaultServeMux。
	port := ":8080"
	fmt.Printf("服务器正在启动，监听端口 %s...\n", port)
	fmt.Printf("请在浏览器中访问: http://localhost%s\n", port)

	// ListenAndServe 会阻塞当前 Goroutine，直到服务器发生错误 (例如端口被占用) 或被关闭。
	// 如果发生错误，它会返回该错误。
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe 错误: %v\n", err) // log.Fatal 会打印错误并退出程序
	}
	// 如果 ListenAndServe 正常返回 (例如服务器被优雅关闭)，程序会继续执行到这里。
	// 但通常它会一直运行，直到你手动停止程序 (Ctrl+C)。
	fmt.Println("服务器已停止。") // 正常情况下这行不会执行
}
