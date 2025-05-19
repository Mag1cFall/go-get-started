package main

import "fmt"

func originalMain() { // <--- 注意这里，main 已被重命名为 originalMain
	// fmt.Println 是一个函数调用，用于在控制台打印一行文本。
	// Println 中的 "ln" 代表 "line"，表示打印后会换行。
	fmt.Println("Hello From VS Code and Git! (from originalMain)")

	// 声明一个整型变量并赋值
	// var 关键字用于声明变量，age 是变量名，int 是类型，18 是初始值。
	var age int = 18
	fmt.Println("我的年龄是:", age)

	// Go 可以自动推断类型，所以你可以省略类型声明
	// name := "Roo" 是短变量声明，等价于 var name string = "Roo"
	// := 只能在函数内部使用
	name := "Roo"
	fmt.Println("我的名字是:", name)

	// 声明一个字符串常量
	// const 关键字用于声明常量，常量的值在编译时确定，不能被修改。
	const greeting string = "你好"
	fmt.Println(greeting, name) // 可以同时打印多个值

	// 条件语句 if-else
	// if 语句的条件不需要用括号括起来
	if age >= 18 {
		fmt.Println("我已经成年了。")
	} else {
		fmt.Println("我还是未成年。")
	}

	// 循环语句 for
	// Go 只有 for 循环，但有多种形式
	// 1. 基本的 for 循环，类似 C 语言的 for
	fmt.Println("基本的 for 循环:")
	for i := 0; i < 3; i++ { // i++ 表示 i = i + 1
		fmt.Println(i)
	}

	// 2. 类似 while 的 for 循环
	fmt.Println("类似 while 的 for 循环:")
	count := 0
	for count < 3 {
		fmt.Println(count)
		count++ // 同样是 count = count + 1
	}

	// 3. 无限循环 (需要 break 来跳出)
	// fmt.Println("无限循环示例 (注释掉了，防止卡住):")
	// for {
	// 	fmt.Println("这是一个无限循环，按 Ctrl+C 停止程序 (如果运行)")
	//  // 通常会有一个条件来 break 跳出循环
	//  break
	// }

	// 调用一个自定义函数
	message := sayHelloOriginal("Go初学者") // <--- sayHello 也重命名以避免与新代码中的函数潜在冲突
	fmt.Println(message)
}

// 自定义一个函数
// func 关键字用于定义函数
// sayHello 是函数名
// (personName string) 是参数列表，personName 是参数名，string 是参数类型
// string 是返回值类型
func sayHelloOriginal(personName string) string { // <--- sayHello 也重命名
	// fmt.Sprintf 用于格式化字符串，但不会打印出来，而是返回格式化后的字符串
	return fmt.Sprintf("你好, %s! 欢迎学习 Go。(from sayHelloOriginal)", personName)
}
