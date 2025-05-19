# Go 学习之旅 (Pure-100%-AI-Generated)

本仓库所有代码示例及注释均由 AI (Google Gemini 2.5 Pro Preview, 模型版本 05-06) 生成，旨在通过“代码示例 + 详细注释”的模式，系统性地辅助初学者学习 Go 语言。

涵盖内容根据提供的学习大纲逐步进行，从基础语法到并发、Web 开发、数据库、测试等。

希望这些示例能为你学习 Go 语言提供有益的参考！

---

## 内容导读

以下是本仓库中按周组织的学习内容和对应的代码示例：

*   **第1周：编程入门与Go语言初探**
    *   核心 Go 语法基础: [`week1/core_syntax/week1_core_syntax.go`](week1/core_syntax/week1_core_syntax.go)

*   **第2周：Go复合类型与结构化编程**
    *   复合类型 (数组, 切片, Map): [`week2/compound_types/week2_compound_types.go`](week2/compound_types/week2_compound_types.go)
    *   指针: [`week2/pointers/week2_pointers.go`](week2/pointers/week2_pointers.go)
    *   结构体: [`week2/structs/week2_structs.go`](week2/structs/week2_structs.go)
    *   方法: [`week2/methods/week2_methods.go`](week2/methods/week2_methods.go)

*   **第3周：接口、包管理与Go Modules**
    *   接口: [`week3/interfaces/week3_interfaces.go`](week3/interfaces/week3_interfaces.go)
    *   包 (自定义包 `geometry` 及使用): [`week3/packages/`](week3/packages/)
    *   Go Modules (添加第三方依赖 `uuid`): [`week3/modules_example/main.go`](week3/modules_example/main.go) (相关变动在根目录的 [`go.mod`](go.mod) 和 [`go.sum`](go.sum))

*   **第4周：错误处理进阶、常用标准库与并发初步**
    *   高级错误处理 (自定义错误, defer, panic, recover): [`week4/advanced_error_handling/week4_advanced_error_handling.go`](week4/advanced_error_handling/week4_advanced_error_handling.go)
    *   标准库 `strings`: [`week4/stdlib_examples/week4_stdlib_strings.go`](week4/stdlib_examples/week4_stdlib_strings.go)
    *   标准库 `strconv`: [`week4/stdlib_examples/week4_stdlib_strconv.go`](week4/stdlib_examples/week4_stdlib_strconv.go)
    *   标准库 `time`: [`week4/stdlib_examples/week4_stdlib_time.go`](week4/stdlib_examples/week4_stdlib_time.go)
    *   标准库 `os` 和 `io` (文件操作): [`week4/stdlib_examples/week4_stdlib_os_io.go`](week4/stdlib_examples/week4_stdlib_os_io.go)
    *   标准库 `encoding/json`: [`week4/stdlib_examples/week4_stdlib_json.go`](week4/stdlib_examples/week4_stdlib_json.go)
    *   并发编程初步 (Goroutines, Channels, WaitGroup): [`week4/concurrency_preliminary/week4_goroutines_channels.go`](week4/concurrency_preliminary/week4_goroutines_channels.go)

*   **第5周：并发编程深入**
    *   Channel深入, `select`语句, `sync`包 (Mutex, RWMutex, Once): [`week5/advanced_concurrency/week5_advanced_concurrency.go`](week5/advanced_concurrency/week5_advanced_concurrency.go)

*   **第6周：网络编程与Web框架初步 (Gin)**
    *   `net/http` 标准库 (简单服务器与客户端): [`week6/net_http_basic/`](week6/net_http_basic/)
    *   Gin 框架入门 (安装, 路由, 参数, 请求数据, 绑定校验, HTML响应, 中间件): [`week6/gin_intro/week6_simple_gin_server.go`](week6/gin_intro/week6_simple_gin_server.go)

*   **第7周：数据库、缓存与原理回顾**
    *   数据库操作 (MySQL 示例框架): [`week7/database_mysql/week7_mysql_example.go`](week7/database_mysql/week7_mysql_example.go)
    *   缓存操作 (Redis 示例框架): [`week7/cache_redis/week7_redis_example.go`](week7/cache_redis/week7_redis_example.go)
    *   核心原理回顾 (make/new, 结构体传递, 反射等): [`week7/core_principles/week7_core_principles.go`](week7/core_principles/week7_core_principles.go)

*   **第8周：项目整合、测试、工具与面试冲刺**
    *   测试 (`testing`包): [`week8/testing_examples/`](week8/testing_examples/) (包含 `math_operations.go` 和 `math_operations_test.go`)
    *   性能分析工具 (`pprof`): [`week8/pprof_example/week8_pprof_server.go`](week8/pprof_example/week8_pprof_server.go)

每个文件都包含了详细的注释，解释了相关的Go语言特性和用法。请按顺序学习，并尝试在本地运行这些示例代码。
