package main

import (
	"fmt"
	// 导入我们自定义的 geometry 包
	// 路径是相对于项目根目录下的 GOPATH/src 或者 Go Modules 的模块路径
	// 在 Go Modules 项目中，如果 geometry 是当前模块的一部分，
	// 导入路径通常是 "moduleName/path/to/package"
	// 例如，如果我们的 go.mod 定义了 module "github.com/Mag1cFall/go-get-started",
	// 那么导入路径就是 "github.com/Mag1cFall/go-get-started/week3/packages/geometry"
	//
	// VS Code 和 Go 工具通常能自动解析同模块下的相对路径，
	// 但标准的导入路径是基于模块的。
	// 为了简单起见，并假设 Go 工具能处理好同模块下的相对引用，我们先用相对路径风格。
	// 如果遇到问题，我们可能需要调整为完整的模块路径。
	//
	// 鉴于当前项目的结构和 go.mod (module github.com/Mag1cFall/go-get-started),
	// 正确的导入路径应该是：
	"github.com/Mag1cFall/go-get-started/week3/packages/geometry"
)

func main() {
	fmt.Println("--- 第3周学习：使用自定义包 ---")

	// 调用 geometry 包中导出的常量
	fmt.Printf("geometry 包中的 Pi 常量: %.4f\n", geometry.Pi)

	// 调用 geometry 包中导出的函数
	rectWidth, rectHeight := 10.0, 5.0
	area := geometry.Area(rectWidth, rectHeight) // 旧的 Area 函数，直接接收参数
	perimeter := geometry.Perimeter(rectWidth, rectHeight)
	fmt.Printf("矩形 (%.2f x %.2f): 面积 = %.2f, 周长 = %.2f\n", rectWidth, rectHeight, area, perimeter)

	fmt.Println("\n--- 使用 geometry 包中的导出类型和方法 ---")
	// 创建 Circle 实例 (Circle 是导出的)
	circ := geometry.Circle{Radius: 7.0}
	fmt.Printf("圆形半径: %.2f\n", circ.Radius)
	fmt.Printf("圆形面积 (通过方法调用): %.2f\n", circ.CircleArea())

	// 创建 Rectangle 实例 (Rectangle 是导出的)
	// 使用 NewRectangle 构造函数
	myRect, err := geometry.NewRectangle(8.0, 4.0)
	if err != nil {
		fmt.Println("创建 Rectangle 失败:", err)
	} else {
		fmt.Printf("自定义矩形: Width=%.2f, Height=%.2f\n", myRect.Width, myRect.Height)
		fmt.Printf("  面积: %.2f\n", myRect.Area())      // 调用 Rectangle 的 Area 方法
		fmt.Printf("  周长: %.2f\n", myRect.Perimeter()) // 调用 Rectangle 的 Perimeter 方法
	}

	// 尝试创建无效的 Rectangle
	_, err = geometry.NewRectangle(-1.0, 5.0)
	if err != nil {
		fmt.Println("创建无效 Rectangle 时捕获到错误:", err)
	}

	// 注意：geometry 包中的 rect 结构体和 internalHelperFunction 函数因为未导出（首字母小写），
	// 所以不能在 main 包中直接访问。
	// var r geometry.rect // 这会导致编译错误: cannot refer to unexported name geometry.rect
	// geometry.internalHelperFunction() // 这会导致编译错误: cannot refer to unexported name geometry.internalHelperFunction

	fmt.Println("\n--- 包的使用演示结束 ---")
	// init 函数的调用顺序：
	// 1. 被导入包的 init 函数 (geometry包的init会先执行)
	// 2. 当前包 (main包) 的 init 函数 (如果定义了的话)
	// 3. main 函数
}

// 我们也可以在 main 包中定义 init 函数
func init() {
	fmt.Println("main 包 (week3/packages/main.go) 的 init 函数被调用了。")
}
