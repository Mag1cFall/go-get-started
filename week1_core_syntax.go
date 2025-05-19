package main

import (
	"errors"
	"fmt"
	"strconv" // 用于字符串转换
)

// main 函数是程序的入口
func main() {
	fmt.Println("--- Go核心语法学习 ---")

	// --- 数据类型 ---
	fmt.Println("\n--- 1. 数据类型 ---")
	// 布尔型 (bool)
	var isLearning bool = true
	var isDifficult bool = false
	fmt.Println("正在学习中?", isLearning)
	fmt.Println("Go语言难吗?", isDifficult)

	// 字符串 (string)
	var courseName string = "Go语言学习"
	fmt.Println("课程名称:", courseName)

	// 数值类型
	// 整型 (int, int8, int16, int32, int64, uint, uintptr 等)
	var apples int = 10
	var oranges int32 = 5
	// Go 会根据操作系统位数推断 int 的具体大小 (32位或64位)
	fmt.Printf("我有 %d 个苹果和 %d 个橘子。\n", apples, oranges) // Printf 用于格式化输出

	// 浮点型 (float32, float64)
	var price float64 = 99.99
	var discount float32 = 0.8
	fmt.Printf("商品价格: %.2f, 折扣: %.1f\n", price, float64(discount)) // %.2f 表示保留两位小数

	// 派生类型 - 举例：类型别名 (后面会学到更多，如数组、切片、map、struct等)
	type MyInteger int
	var myApples MyInteger = 12
	fmt.Println("我的苹果 (自定义类型):", myApples)

	// --- fmt 包的更多用法 ---
	fmt.Println("\n--- 2. fmt 包 ---")
	name := "Roo"
	age := 3
	// Printf: 格式化输出
	fmt.Printf("大家好，我是 %s，今年 %d 岁。\n", name, age)
	// Sprintf: 格式化并返回字符串，不打印
	userInfo := fmt.Sprintf("用户信息: 姓名-%s, 年龄-%d", name, age)
	fmt.Println(userInfo)

	// Scanln: 从标准输入读取 (为了不阻塞自动流程，这里注释掉，你可以取消注释自行测试)
	// fmt.Println("请输入你的名字:")
	// var inputName string
	// fmt.Scanln(&inputName) // & 取变量地址
	// fmt.Printf("你好, %s!\n", inputName)

	// --- 流程控制：switch 语句 ---
	fmt.Println("\n--- 3. switch 语句 ---")
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("开始新的一周！")
	case "星期三":
		fmt.Println("周中加油！")
		// Go的switch默认带break，不需要显式写
	case "星期五":
		fmt.Println("周末快到了！")
	default:
		fmt.Println("平常的一天。")
	}

	// switch 还可以不带表达式，作为 if/else if 的替代
	score := 85
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// --- 函数：多返回值与错误处理 ---
	fmt.Println("\n--- 4. 函数：多返回值与错误处理 ---")
	result, err := divide(10, 2)
	if err != nil {
		// 通常错误处理是打印日志或返回错误
		fmt.Println("除法错误:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0) // 尝试除以0
	if err != nil {
		fmt.Println("除法错误:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}

	// --- 函数：匿名函数与闭包初步 ---
	fmt.Println("\n--- 5. 函数：匿名函数与闭包初步 ---")
	// 匿名函数：没有名字的函数，可以直接赋值给变量或直接执行
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("匿名函数计算 5 + 3 =", add(5, 3))

	// 立即执行的匿名函数
	func(message string) {
		fmt.Println("立即执行:", message)
	}("你好呀！")

	// 闭包：一个函数"记住"了其外部作用域的变量
	// (更深入的闭包会在后续学习中体现)
	counter := createCounter()
	fmt.Println(counter()) // 输出 1
	fmt.Println(counter()) // 输出 2
	fmt.Println(counter()) // 输出 3

	newCounter := createCounter()
	fmt.Println(newCounter()) // 输出 1 (这是一个新的计数器实例)

	// 调用字符串转换示例函数
	stringConversionExample()

	fmt.Println("\n--- 核心语法学习告一段落 ---")
}

// divide 函数演示多返回值和错误处理
// (num1 int, num2 int) 是参数列表
// (int, error) 是返回值列表，第一个是结果，第二个是错误对象
func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		// errors.New 创建一个新的错误对象
		return 0, errors.New("除数不能为零")
	}
	return num1 / num2, nil // nil 表示没有错误
}

// createCounter 函数演示闭包
// 它返回一个函数，这个返回的函数可以访问和修改 createCounter 作用域内的变量 (i)
func createCounter() func() int {
	i := 0 // 这个 i 对于返回的匿名函数是可见的
	return func() int {
		i++
		return i
	}
}

// 补充：strconv 包用于字符串和其他类型的转换
func stringConversionExample() {
	fmt.Println("\n--- strconv 包示例 ---")
	s := "123"
	// Atoi: string to int (ASCII to Integer)
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("字符串转整数失败:", err)
	} else {
		fmt.Println(s, "转换为整数是:", num)
	}

	n := 456
	// Itoa: int to string (Integer to ASCII)
	str := strconv.Itoa(n)
	fmt.Println(n, "转换为字符串是:", str)

	boolStr := "true"
	// ParseBool: string to bool
	b, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("字符串转布尔失败:", err)
	} else {
		fmt.Println(boolStr, "转换为布尔是:", b)
	}
	// 注意：stringConversionExample 函数没有在 main 中被调用，
	// 如果需要运行它，可以在 main 函数中添加 stringConversionExample()
}
