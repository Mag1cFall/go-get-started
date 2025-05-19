package main

import (
	"fmt"
	"io"
	"net/http" // 导入 net/http 包
	"net/url"  // 导入 net/url 包，用于 POST 表单数据
	"os"
	"strings"
)

// helper function to read and print response body
func printResponseBody(res *http.Response, actionDesc string) {
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("  [Client] 读取 %s 响应体错误: %v\n", actionDesc, err)
		return
	}
	fmt.Printf("  [Client] %s 响应状态码: %s\n", actionDesc, res.Status)
	fmt.Printf("  [Client] %s 响应体:\n%s\n", actionDesc, string(bodyBytes))
}

func main() {
	fmt.Println("--- 第6周学习：net/http 标准库 (简单 HTTP 客户端) ---")
	// 确保服务器 (week6_simple_server.go) 正在另一个终端的同一目录下运行，监听 :8080

	serverURL := "http://localhost:8080"

	// --- 1. 发送 GET 请求 ---
	fmt.Println("\n--- 1. 发送 GET 请求 ---")

	// a) GET /hello
	fmt.Println("  [Client] 正在发送 GET 请求到:", serverURL+"/hello")
	resHello, err := http.Get(serverURL + "/hello")
	if err != nil {
		fmt.Printf("  [Client] GET /hello 错误: %v\n", err)
		fmt.Println("  请确保 week6_simple_server.go 正在运行!")
		os.Exit(1) // 如果服务器没运行，后续请求也无法成功，直接退出
	}
	defer resHello.Body.Close() // 非常重要：确保关闭响应体以释放资源
	printResponseBody(resHello, "GET /hello")

	// b) GET /time
	fmt.Println("\n  [Client] 正在发送 GET 请求到:", serverURL+"/time")
	resTime, err := http.Get(serverURL + "/time")
	if err != nil {
		fmt.Printf("  [Client] GET /time 错误: %v\n", err)
	} else {
		defer resTime.Body.Close()
		printResponseBody(resTime, "GET /time")
	}

	// c) GET /headers (这个端点在服务器端会返回客户端发送的请求头)
	fmt.Println("\n  [Client] 正在发送 GET 请求到:", serverURL+"/headers")
	// 我们可以创建一个自定义的请求来添加请求头
	reqHeaders, err := http.NewRequest("GET", serverURL+"/headers", nil)
	if err != nil {
		fmt.Printf("  [Client] 创建 GET /headers 请求错误: %v\n", err)
	} else {
		reqHeaders.Header.Add("X-Custom-Header", "GoClientTest")
		reqHeaders.Header.Add("User-Agent", "MyGoClient/1.0")

		client := &http.Client{}                 // 创建一个 HTTP 客户端
		resHeaders, err := client.Do(reqHeaders) // 发送请求
		if err != nil {
			fmt.Printf("  [Client] 发送 GET /headers 请求错误: %v\n", err)
		} else {
			defer resHeaders.Body.Close()
			printResponseBody(resHeaders, "GET /headers")
		}
	}

	// --- 2. 发送 POST 请求 ---
	// 我们没有在 simple_server 中定义处理 POST 请求的端点，
	// 但我们可以演示如何发送一个 POST 请求。
	// 如果发送到服务器的根路径 "/"，它可能会被默认处理器处理。
	fmt.Println("\n--- 2. 发送 POST 请求 (示例) ---")

	// a) POST application/x-www-form-urlencoded
	formData := url.Values{}
	formData.Set("name", "Go Developer")
	formData.Set("project", "HTTP Client Example")

	fmt.Println("  [Client] 正在发送 POST (form-urlencoded) 请求到:", serverURL+"/")
	// http.PostForm 发送 Content-Type: application/x-www-form-urlencoded
	resPostForm, err := http.PostForm(serverURL+"/", formData)
	if err != nil {
		fmt.Printf("  [Client] POST / (form) 错误: %v\n", err)
	} else {
		defer resPostForm.Body.Close()
		printResponseBody(resPostForm, "POST / (form-urlencoded)")
	}

	// b) POST application/json
	// (我们的简单服务器没有专门处理JSON的端点，但可以演示发送)
	jsonBody := `{"message":"Hello from JSON POST","value":123}`
	fmt.Println("  [Client] 正在发送 POST (json) 请求到:", serverURL+"/hello") // 发到 /hello 看看服务器怎么响应
	// http.Post 发送指定 Content-Type 的 POST 请求
	resPostJSON, err := http.Post(serverURL+"/hello", "application/json", strings.NewReader(jsonBody))
	if err != nil {
		fmt.Printf("  [Client] POST /hello (json) 错误: %v\n", err)
	} else {
		defer resPostJSON.Body.Close()
		printResponseBody(resPostJSON, "POST /hello (json)")
	}

	fmt.Println("\n--- HTTP 客户端演示结束 ---")
}
