package main

import (
	"fmt"
	"net/http" // 导入 net/http 包，Gin 内部使用它，并且我们也用它来定义状态码
	"time"     // <--- 添加 time 包的导入，用于中间件

	"github.com/gin-gonic/gin" // 导入 Gin 包
)

// SimpleLoggerMiddleware 是一个简单的自定义日志中间件
// 中间件本质上是一个 gin.HandlerFunc
func SimpleLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// 在请求处理之前可以做一些事情
		fmt.Printf("[自定义中间件] 请求开始: %s %s\n", c.Request.Method, path)

		// 调用 c.Next() 来执行后续的处理器 (包括其他中间件和路由处理器)
		c.Next()

		// 在请求处理之后可以做一些事情
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		fmt.Printf("[自定义中间件] 请求完成: %s | %3d | %13v | %s \n",
			path,
			statusCode,
			latency,
			c.ClientIP(),
		)
	}
}

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

	// --- 5. 获取请求数据 (Form表单, JSON Body) 与参数绑定校验 ---
	// a) 处理 POST Form 表单数据
	// 客户端可以用 application/x-www-form-urlencoded 或 multipart/form-data 发送
	router.POST("/form_post", func(c *gin.Context) {
		// c.PostForm(key string) 获取表单字段的值
		message := c.PostForm("message")
		// c.DefaultPostForm(key, defaultValue string) 获取表单字段，若不存在则用默认值
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// b) 处理 POST JSON Body 数据并绑定到结构体
	type LoginPayload struct {
		// `binding:"required"` 标签表示该字段是必需的
		// Gin 使用 `github.com/go-playground/validator/v10` 进行校验
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"` // 密码至少6位
		Email    string `json:"email" binding:"omitempty,email"`   // 可选，但如果是，则必须是email格式
	}

	router.POST("/login_json", func(c *gin.Context) {
		var loginData LoginPayload

		// c.ShouldBindJSON(&obj) 会尝试将请求的 JSON body 绑定到 loginData 结构体。
		// 如果绑定失败或校验失败 (基于 struct tags)，会返回错误。
		if err := c.ShouldBindJSON(&loginData); err != nil {
			// 如果校验失败，返回 400 Bad Request 错误，并附带错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 校验通过，处理登录逻辑 (此处仅为演示)
		if loginData.Username == "testuser" && loginData.Password == "123456" {
			c.JSON(http.StatusOK, gin.H{
				"status":   "login successful",
				"username": loginData.Username,
				"email":    loginData.Email, // Email 会被正确绑定，即使它是可选的
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "login failed", "message": "Invalid credentials"})
		}
	})

	// --- 6. 返回 HTML 响应 ---
	// Gin 可以加载和渲染 HTML 模板。为了简单起见，这里我们直接返回 HTML 字符串。
	// 更复杂的场景会使用 router.LoadHTMLGlob("templates/*") 或 router.LoadHTMLFiles("template1.html", "template2.html")
	// 然后在处理器中使用 c.HTML(http.StatusOK, "templateName.html", data)
	router.GET("/html_page", func(c *gin.Context) {
		htmlContent := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Gin HTML Page</title>
			</head>
			<body>
				<h1>你好，来自 Gin 的 HTML 页面!</h1>
				<p>这是一个通过 c.Data() 直接返回的简单 HTML 示例。</p>
			</body>
			</html>
		`
		// c.Data(statusCode, contentType, data []byte)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})

	// --- 7. 中间件 (Middleware) ---
	// Gin 引擎默认使用了 Logger 和 Recovery 中间件。
	// Logger: 记录每个请求的日志到控制台。
	// Recovery: 捕获任何 panic 并返回 500 错误，防止服务器崩溃。

	// a) 使用我们定义在包级别的 SimpleLoggerMiddleware
	// router.Use(middleware ...gin.HandlerFunc) 可以注册全局中间件，对所有路由生效。
	router.Use(SimpleLoggerMiddleware()) // <--- 使用移到包级别的中间件

	// b) 也可以为特定的路由或路由组注册中间件
	adminGroup := router.Group("/admin")
	adminGroup.Use(func(c *gin.Context) { // 另一个简单的内联中间件
		fmt.Println("[Admin中间件] 检查管理员权限...")
		// 假设这里有一些权限检查逻辑
		// if !isAdmin { c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error":"无权限"}); return }
		c.Next()
		fmt.Println("[Admin中间件] 管理员权限检查通过。")
	})
	{
		adminGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "欢迎来到管理员面板!"})
		})
	}

	// --- 8. 启动 Gin 服务器 ---
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
	fmt.Printf("  POST http://localhost%s/form_post (Form表单)\n", port)
	fmt.Printf("  POST http://localhost%s/login_json (JSON Body与绑定校验)\n", port)
	fmt.Printf("  http://localhost%s/html_page (HTML响应)\n", port)
	fmt.Printf("  http://localhost%s/admin/dashboard (带中间件的路由组)\n", port)

	// Run 会阻塞当前 Goroutine，直到服务器发生错误或被关闭。
	// 如果发生错误，它会 panic。
	err := router.Run(port)
	if err != nil {
		// 通常 Run() 发生错误会直接 panic，所以这行可能不会执行到，
		// 除非是某些特定类型的错误。Gin 的 Recovery 中间件会处理 panic。
		fmt.Printf("Gin 服务器启动失败: %v\n", err)
	}
}
