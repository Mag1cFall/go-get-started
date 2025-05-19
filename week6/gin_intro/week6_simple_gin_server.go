package main

import (
	"fmt"
	"net/http" // 导入 net/http 包，Gin 内部使用它，并且我们也用它来定义状态码

	"github.com/gin-gonic/gin" // 导入 Gin 包
)

func main() {
	fmt.Println("--- 第6周学习：Gin 框架入门 (简单服务器与路由) ---")

	// --- 1. 创建 Gin 引擎 ---
	// gin.Default() 返回一个默认的 Gin 引擎实例。
	// 它已经包含了 Logger (日志) 和 Recovery (恐慌恢复) 中间件。
	// 如果你想要一个没有任何中间件的纯净引擎，可以使用 gin.New()。
	router := gin.Default()

	// --- 2. 定义处理器函数 (Handler Functions for Gin) ---
	// Gin 的处理器函数接收一个 *gin.Context 参数。
	// gin.Context 封装了 HTTP 请求和响应，并提供了很多便捷的方法。

	// pingHandler 响应 "/ping" 路径的 GET 请求
	pingHandler := func(c *gin.Context) {
		// c.JSON 用于发送 JSON 格式的响应。
		// 第一个参数是 HTTP 状态码，第二个参数是要序列化为 JSON 的数据 (可以是 map, struct 等)。
		// gin.H 是 map[string]interface{} 的一个快捷方式。
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from Gin!",
			"status":  "success",
		})
	}

	// --- 3. 注册路由和处理器 ---
	// 使用 Gin 引擎的 HTTP 方法函数 (如 GET, POST, PUT, DELETE 等) 来注册路由。
	// 第一个参数是路由路径，后续参数是处理器函数(可以有多个，形成处理器链)。

	// GET 请求到 "/ping"
	router.GET("/ping", pingHandler)

	// GET 请求到根路径 "/"
	router.GET("/", func(c *gin.Context) {
		// c.String 用于发送纯文本响应。
		c.String(http.StatusOK, "欢迎来到我的第一个 Gin 服务器!")
	})

	// GET 请求带路径参数
	// 路径参数以冒号 : 开头，例如 :name
	router.GET("/user/:name", func(c *gin.Context) { // 路径修改为 /user/:name 以区分
		// c.Param(key string) 用于获取路径参数的值。
		name := c.Param("name")
		message := fmt.Sprintf("你好用户, %s!", name)
		c.JSON(http.StatusOK, gin.H{"user_message": message}) // 返回JSON
	})

	// 获取查询参数 (Query Parameters)
	// 例如: /search?query=golang&sort=asc
	router.GET("/search", func(c *gin.Context) {
		// c.Query(key string) 获取指定名称的查询参数。如果不存在，返回空字符串。
		query := c.Query("query")
		// c.DefaultQuery(key, defaultValue string) 获取查询参数，如果不存在，则使用指定的默认值。
		sortOrder := c.DefaultQuery("sort", "desc") // 默认降序

		c.JSON(http.StatusOK, gin.H{
			"search_term": query,
			"sorted_by":   sortOrder,
			"results":     []string{fmt.Sprintf("结果1 for %s", query), fmt.Sprintf("结果2 for %s", query)},
		})
	})

	// --- 4. 路由组 (Route Grouping) ---
	// 路由组允许我们将具有共同前缀或共同中间件的路由组织在一起。
	// 例如，所有 /api/v1/... 的路由可以放在一个组里。
	apiV1 := router.Group("/api/v1")
	{ // 这个花括号只是为了视觉上的组织，不是必需的
		// /api/v1/users
		apiV1.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"group":    "v1",
				"resource": "users",
				"data":     []string{"userA", "userB", "userC"},
			})
		})

		// /api/v1/products
		apiV1.GET("/products", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"group":    "v1",
				"resource": "products",
				"data":     []string{"productX", "productY"},
			})
		})
	} // 结束 apiV1 组

	// --- 5. 启动 Gin 服务器 ---
	// router.Run() 启动 HTTP 服务器，并监听指定的地址和端口。
	// 如果不提供参数，默认监听 ":8080"。
	// 你也可以指定其他端口，例如 router.Run(":8888")。
	port := ":8080" // 和我们之前的 net/http 服务器使用相同端口，确保之前的已停止
	fmt.Printf("Gin 服务器正在启动，监听端口 %s...\n", port)
	fmt.Printf("请在浏览器或工具中访问:\n")
	fmt.Printf("  http://localhost%s/\n", port)
	fmt.Printf("  http://localhost%s/ping\n", port)
	fmt.Printf("  http://localhost%s/user/你的名字 (路径参数)\n", port)
	fmt.Printf("  http://localhost%s/search?query=Go&sort=asc (查询参数)\n", port)
	fmt.Printf("  http://localhost%s/api/v1/users (路由组)\n", port)
	fmt.Printf("  http://localhost%s/api/v1/products (路由组)\n", port)

	// Run 会阻塞当前 Goroutine，直到服务器发生错误或被关闭。
	// 如果发生错误，它会 panic。
	err := router.Run(port)
	if err != nil {
		// 通常 Run() 发生错误会直接 panic，所以这行可能不会执行到，
		// 除非是某些特定类型的错误。Gin 的 Recovery 中间件会处理 panic。
		fmt.Printf("Gin 服务器启动失败: %v\n", err)
	}
}
