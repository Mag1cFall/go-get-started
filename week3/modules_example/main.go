package main

import (
	"fmt"
	// 导入第三方包
	"github.com/google/uuid"
)

func main() {
	fmt.Println("--- 第3周学习：Go Modules 与第三方依赖 ---")

	// 生成一个新的 UUID
	newUUID, err := uuid.NewRandom()
	if err != nil {
		// 处理错误，例如记录日志或退出
		// 在实际应用中，错误处理会更复杂
		fmt.Printf("生成 UUID 失败: %v\n", err)
		return
	}

	fmt.Printf("生成的 UUID: %s\n", newUUID.String())

	// 另一个例子：从字符串解析 UUID
	// 这是一个有效的 UUID v4 字符串示例
	uuidStr := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		fmt.Printf("解析 UUID 字符串 '%s' 失败: %v\n", uuidStr, err)
	} else {
		fmt.Printf("从字符串解析的 UUID: %s\n", parsedUUID.String())
		fmt.Printf("  版本: %s\n", parsedUUID.Version().String())
		fmt.Printf("  变体: %s\n", parsedUUID.Variant().String())
	}

	fmt.Println("\n--- Go Modules 演示结束 ---")
	// 当你运行这个文件时 (例如 go run main.go)，Go 工具会自动执行以下操作：
	// 1. 检查 go.mod 文件中是否已列出 github.com/google/uuid 这个依赖。
	// 2. 如果没有，它会去下载这个包的最新版本。
	// 3. 更新 go.mod 文件，将这个新的依赖及其版本添加进去。
	// 4. 创建或更新 go.sum 文件，记录这个依赖包及其所有传递依赖包的校验和，以确保依赖的完整性和一致性。
}
