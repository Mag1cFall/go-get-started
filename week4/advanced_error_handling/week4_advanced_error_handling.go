package main

import (
	"fmt"
	"os" // 用于 defer 示例中的文件操作
)

// --- 1. 自定义错误类型 ---
// Go中，错误类型是任何实现了 error 接口的类型。
// type error interface {
//    Error() string
// }
// 通过实现这个接口，我们可以创建具有更多上下文信息的自定义错误类型。

// MyError 是一个自定义错误类型
type MyError struct {
	Operation string // 进行的操作
	Message   string // 错误信息
	ErrorCode int    // 错误码
}

// 实现 error 接口的 Error() string 方法
func (e *MyError) Error() string {
	return fmt.Sprintf("操作 '%s' 失败: %s (错误码: %d)", e.Operation, e.Message, e.ErrorCode)
}

// 一个可能返回自定义错误的函数
func performSensitiveOperation(shouldFail bool) error {
	if shouldFail {
		return &MyError{ // 返回自定义错误类型的指针
			Operation: "敏感数据处理",
			Message:   "权限不足或数据损坏",
			ErrorCode: 1001,
		}
	}
	fmt.Println("敏感操作成功完成。")
	return nil
}

// --- 2. defer 语句 ---
// defer 语句会将其后面跟随的函数调用（称为延迟函数）推迟到紧邻的包含 defer 语句的函数返回之前执行。
// 延迟调用的参数会立即求值，但函数调用本身会推迟。
// 如果有多个 defer 语句，它们会以“后进先出”（LIFO）的顺序执行。
// defer 常用于确保资源（如文件、网络连接、锁）在函数结束时被释放。

func deferExample() {
	fmt.Println("  deferExample: 开始")

	defer fmt.Println("  deferExample: 第一个 defer (最后执行)") // 3
	defer fmt.Println("  deferExample: 第二个 defer (中间执行)") // 2

	fmt.Println("  deferExample: 函数体执行中...")

	defer fmt.Println("  deferExample: 第三个 defer (最先执行)") // 1

	fmt.Println("  deferExample: 结束")
	// 返回前，会按 LIFO 顺序执行 defer 后的函数调用：
	// 1. "  deferExample: 第三个 defer (最先执行)"
	// 2. "  deferExample: 第二个 defer (中间执行)"
	// 3. "  deferExample: 第一个 defer (最后执行)"
}

func fileOperationWithDefer() {
	fmt.Println("  fileOperationWithDefer: 尝试打开文件...")
	file, err := os.Create("temp.txt") // 尝试创建一个临时文件
	if err != nil {
		fmt.Println("  创建文件失败:", err)
		return
	}
	// 使用 defer 确保文件在函数退出前关闭，无论函数如何退出（正常返回或panic）
	defer file.Close()
	defer fmt.Println("  fileOperationWithDefer: 文件关闭操作已注册 (defer file.Close())")

	fmt.Println("  fileOperationWithDefer: 文件创建成功，写入数据...")
	_, err = file.WriteString("Hello from defer example!")
	if err != nil {
		fmt.Println("  写入文件失败:", err)
		// file.Close() 会在这里由 defer 调用
		return
	}
	fmt.Println("  fileOperationWithDefer: 数据写入成功。")
	// file.Close() 会在这里由 defer 调用
}

// --- 3. panic 和 recover ---
// panic: 当程序遇到无法处理的严重错误时（例如数组越界、空指针解引用），会触发 panic。
//        panic 会立即停止当前函数的执行，并开始逐层向上执行包含它的函数的 defer 语句，
//        然后程序会崩溃并打印错误信息和堆栈跟踪。
//        也可以通过调用内置的 panic() 函数主动触发。
//
// recover: 是一个内置函数，用于重新获得对一个已经发生 panic 的 goroutine 的控制权。
//          recover 只有在 defer 函数内部被直接调用时才有效。
//          如果当前 goroutine 没有发生 panic，或者 recover 不是在 defer 中被调用，它会返回 nil。
//          如果发生了 panic，recover 会捕获到传递给 panic 的值，并允许程序从 panic 中恢复，
//          阻止程序崩溃（但当前函数的执行仍然会停止）。
//
// 通常，不应滥用 panic 和 recover 来进行正常的错误处理。
// 它们主要用于处理真正的意外或不可恢复的错误，或者在库代码的边界防止内部 panic 泄露给调用者。

func mightPanic(shouldPanic bool) {
	defer fmt.Println("  mightPanic: defer 语句执行 (在 panic 发生后，或正常返回前)")

	if shouldPanic {
		fmt.Println("  mightPanic: 准备触发 panic!")
		panic("这是一个故意的 panic!") // 主动触发 panic
		// panic 之后的代码不会执行
		// fmt.Println("这行代码不会被执行")
	}
	fmt.Println("  mightPanic: 函数正常结束。")
}

// safeCall演示了如何使用 recover 来捕获 panic
func safeCall(fn func()) {
	defer func() {
		// recover() 必须在 defer 函数中直接调用
		if r := recover(); r != nil {
			// r 是传递给 panic() 的值
			fmt.Printf("  safeCall: 捕获到 panic: %v\n", r)
			fmt.Println("  safeCall: 程序从 panic 中恢复，不会崩溃。")
		} else {
			fmt.Println("  safeCall: 函数正常执行完毕，没有 panic 发生。")
		}
	}() // 注意这里的 ()，立即执行这个匿名 defer 函数

	fmt.Println("  safeCall: 准备调用函数...")
	fn() // 调用传入的函数，这个函数可能会 panic
	fmt.Println("  safeCall: 函数调用完成 (如果未发生 panic)。")
}

func main() {
	fmt.Println("--- 第4周学习：错误处理进阶 (自定义错误, defer, panic, recover) ---")

	// --- 自定义错误 ---
	fmt.Println("\n--- 1. 自定义错误类型 ---")
	err := performSensitiveOperation(true) // 模拟操作失败
	if err != nil {
		fmt.Println("操作出错:", err) // 会调用 MyError 的 Error() 方法

		// 可以使用类型断言来检查错误的具体类型并访问其字段
		if myErr, ok := err.(*MyError); ok { // 注意是指针类型 *MyError
			fmt.Printf("  这是一个 MyError 类型的错误。操作: %s, 错误码: %d\n",
				myErr.Operation, myErr.ErrorCode)
		}
	}

	fmt.Println()
	err = performSensitiveOperation(false) // 模拟操作成功
	if err != nil {
		fmt.Println("操作出错:", err)
	}

	// --- defer ---
	fmt.Println("\n--- 2. defer 语句 ---")
	deferExample()
	fmt.Println()
	fileOperationWithDefer()
	// 清理临时文件 (在实际应用中，你可能想在测试后删除它)
	// os.Remove("temp.txt") // 可以在这里删除，或者让用户手动删除

	// --- panic 和 recover ---
	fmt.Println("\n--- 3. panic 和 recover ---")

	fmt.Println("\n调用 safeCall 执行一个不会 panic 的函数:")
	safeCall(func() {
		fmt.Println("    这是一个安全的函数调用，不会 panic。")
	})

	fmt.Println("\n调用 safeCall 执行一个会 panic 的函数:")
	safeCall(func() {
		fmt.Println("    准备在函数内部调用 mightPanic(true)...")
		mightPanic(true) // 这个函数会 panic
		fmt.Println("    mightPanic(true) 调用之后 (这行不会执行，因为 panic 了)")
	})

	fmt.Println("\n在 safeCall 之外直接调用会 panic 的函数 (会导致程序崩溃):")
	// 为了防止整个学习流程中断，我们将下面这行注释掉。
	// 如果取消注释，程序会在这里因为未捕获的 panic 而终止。
	// mightPanic(true)
	fmt.Println("  (上面会 panic 的调用已被注释)")

	fmt.Println("\n--- 错误处理进阶学习结束 ---")
}
