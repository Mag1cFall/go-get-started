package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包，下划线表示只执行其init函数以注册HTTP处理器
	"runtime"
	"strings"
	"time"
)

// a computationally intensive function to generate CPU load
func cpuIntensiveTask() {
	for i := 0; i < 100000000; i++ {
		_ = strings.Contains("abcdefghijklmnopqrstuvwxyz", "m")
	}
}

// a function that allocates memory
func memoryAllocatingTask() {
	// Allocate a large slice many times to simulate memory pressure
	for i := 0; i < 50; i++ {
		_ = make([]byte, 1024*1024) // Allocate 1MB
		time.Sleep(50 * time.Millisecond)
	}
}

// a function that creates many goroutines
func createGoroutines() {
	for i := 0; i < 50; i++ {
		go func() {
			time.Sleep(30 * time.Second) // Keep goroutines alive for a while
		}()
	}
	fmt.Printf("  [Server] 当前 Goroutine 数量: %d\n", runtime.NumGoroutine())
}

func main() {
	fmt.Println("--- 第8周学习：性能分析工具 (pprof) ---")

	// 注册一些普通的 HTTP 处理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎来到 pprof 演示服务器! 访问 /debug/pprof/ 来查看性能数据。\n")
		fmt.Fprintf(w, "你可以尝试访问:\n")
		fmt.Fprintf(w, "  /loadcpu - 触发一些 CPU 负载\n")
		fmt.Fprintf(w, "  /allocmem - 触发一些内存分配\n")
		fmt.Fprintf(w, "  /creategoroutines - 创建一些 goroutines\n")
	})

	http.HandleFunc("/loadcpu", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  [Server] 开始执行 CPU 密集型任务...")
		go cpuIntensiveTask() // 在 goroutine 中执行，避免阻塞主线程太久
		fmt.Fprintf(w, "CPU 密集型任务已在后台启动。请稍后通过 pprof 查看 CPU profile。\n")
		fmt.Println("  [Server] CPU 密集型任务已提交。")
	})

	http.HandleFunc("/allocmem", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  [Server] 开始执行内存分配任务...")
		go memoryAllocatingTask() // 在 goroutine 中执行
		fmt.Fprintf(w, "内存分配任务已在后台启动。请稍后通过 pprof 查看 heap profile。\n")
		fmt.Println("  [Server] 内存分配任务已提交。")
	})

	http.HandleFunc("/creategoroutines", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  [Server] 开始创建 Goroutines...")
		createGoroutines()
		fmt.Fprintf(w, "已创建多个 Goroutines。请稍后通过 pprof 查看 goroutine profile。\n")
		fmt.Println("  [Server] Goroutines 创建任务已提交。")
	})

	// pprof 的 HTTP 端点会自动注册在 /debug/pprof/ 路径下，
	// 因为我们导入了 _ "net/http/pprof"。

	port := ":8081" // 使用一个与之前不同的端口，以避免冲突
	fmt.Printf("pprof 演示服务器正在启动，监听端口 %s...\n", port)
	fmt.Printf("请在浏览器中访问: http://localhost%s/debug/pprof/\n", port)
	fmt.Println("服务器启动后，你可以让它运行一段时间，并访问上面的 /loadcpu, /allocmem, /creategoroutines 端点来产生一些负载。")
	fmt.Println("然后，可以使用 'go tool pprof http://localhost:8081/debug/pprof/profile?seconds=30' (CPU) 或")
	fmt.Println("           'go tool pprof http://localhost:8081/debug/pprof/heap' (内存) 等命令进行分析。")

	err := http.ListenAndServe(port, nil) // 使用默认的 ServeMux，pprof 处理器已注册到其中
	if err != nil {
		log.Fatalf("ListenAndServe 错误: %v\n", err)
	}
}

// 辅助函数，用于生成随机字符串（如果需要模拟更复杂的CPU负载）
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
