package main

import "fmt"

func main() {
	fmt.Println("--- 第2周学习：指针 (Pointers) ---")

	// --- 1. 什么是变量和内存地址 ---
	fmt.Println("\n--- 1. 变量和内存地址 ---")
	x := 100
	fmt.Println("变量 x 的值:", x)
	// 使用 & 操作符获取变量的内存地址
	fmt.Println("变量 x 的内存地址:", &x) // 输出会是一个十六进制的地址，例如 0xc0000a2008

	// --- 2. 什么是指针 ---
	// 指针是一个变量，其值为另一个变量的内存地址。
	// 指针变量声明：var pointerName *Type
	fmt.Println("\n--- 2. 指针的声明与初始化 ---")
	var p *int                   // p 是一个指向 int 类型的指针，当前为 nil (零值)
	fmt.Println("未初始化的指针 p:", p) // 输出: <nil>

	if p == nil {
		fmt.Println("指针 p 是 nil, 不能解引用。")
	}

	// 将变量 x 的地址赋值给指针 p
	p = &x
	fmt.Println("指针 p 指向的地址:", p)  // 输出与 &x 相同
	fmt.Println("变量 x 的内存地址:", &x) // 再次确认

	// --- 3. 解引用指针 ---
	// 使用 * 操作符（解引用操作符）来访问指针所指向地址的变量的值。
	fmt.Println("\n--- 3. 解引用指针 ---")
	if p != nil {
		fmt.Println("通过指针 p 获取 x 的值 (*p):", *p) // 输出: 100

		// 通过指针修改其所指向变量的值
		*p = 200                          // 等价于 x = 200
		fmt.Println("通过指针修改后，x 的值:", x)   // 输出: 200
		fmt.Println("通过指针修改后，*p 的值:", *p) // 输出: 200
	}

	// --- 4. 指针的用途 ---
	fmt.Println("\n--- 4. 指针的用途 ---")

	// a) 在函数间共享和修改数据
	// Go 函数参数默认是值传递。如果想在函数内部修改外部变量的值，可以使用指针。
	num := 50
	fmt.Println("调用 modifyValueByVal 前, num:", num)
	modifyValueByVal(num)                           // 值传递，num 的副本被修改
	fmt.Println("调用 modifyValueByVal 后, num:", num) // num 仍然是 50

	fmt.Println("调用 modifyValueByPtr 前, num:", num)
	modifyValueByPtr(&num)                          // 地址传递 (传递指针)
	fmt.Println("调用 modifyValueByPtr 后, num:", num) // num 变为 500

	// b) 提高性能 (对于大型数据结构)
	// 当传递大型结构体时，传递指针比传递整个结构体的副本更高效，因为它只复制地址。
	// (我们将在学习结构体时更详细地看到这一点)

	// c) 表示可选值或可变状态
	// 指针可以是 nil，这可以用来表示一个值不存在或未初始化。

	// --- 5. 指针的指针 (多级指针) ---
	fmt.Println("\n--- 5. 指针的指针 ---")
	a := 10
	var ptrA *int = &a
	var ptrPtrA **int = &ptrA // 指向指针的指针

	fmt.Println("a 的值:", a)
	fmt.Println("ptrA 指向的地址:", ptrA, "ptrA 存储的值 (a的地址):", ptrA)
	fmt.Println("ptrPtrA 指向的地址 (ptrA的地址):", ptrPtrA)

	fmt.Println("*ptrA (a的值):", *ptrA)
	fmt.Println("**ptrPtrA (a的值):", **ptrPtrA)

	**ptrPtrA = 11 // 修改 a 的值
	fmt.Println("通过 **ptrPtrA 修改后, a 的值:", a)

	// --- 6. 不要对 nil 指针解引用 ---
	fmt.Println("\n--- 6. nil 指针 ---")
	var nilPtr *int
	fmt.Println("nilPtr 的值:", nilPtr)
	// *nilPtr = 10 // 这行代码如果取消注释并运行，会导致 panic: runtime error: invalid memory address or nil pointer dereference
	if nilPtr != nil {
		fmt.Println("nilPtr 指向的值:", *nilPtr)
	} else {
		fmt.Println("nilPtr 是 nil，不能安全解引用。")
	}

	// --- 7. new() 函数创建指针 ---
	// new(T) 函数会为类型 T 的新项分配空间，并返回其地址，即一个 *T 类型的值。
	// 这个新项会被初始化为其类型的零值。
	fmt.Println("\n--- 7. new() 函数 ---")
	ptrUsingNew := new(int) // ptrUsingNew 是一个 *int 类型，指向一个值为 0 的 int
	fmt.Println("使用 new 创建的指针 ptrUsingNew:", ptrUsingNew)
	fmt.Println("ptrUsingNew 指向的值 (*ptrUsingNew):", *ptrUsingNew) // 输出: 0
	*ptrUsingNew = 42
	fmt.Println("修改后, *ptrUsingNew:", *ptrUsingNew)

	strPtr := new(string) // 指向一个空字符串 ""
	fmt.Println("使用 new 创建的字符串指针 strPtr:", strPtr)
	fmt.Println("*strPtr:", *strPtr) // 输出: "" (空字符串)
	*strPtr = "Hello from new pointer"
	fmt.Println("*strPtr:", *strPtr)

	fmt.Println("\n--- 指针学习结束 ---")
}

// modifyValueByVal 接收一个 int 值的副本
func modifyValueByVal(val int) {
	val = val * 10 // 修改的是副本
	fmt.Println("在 modifyValueByVal内部, val:", val)
}

// modifyValueByPtr 接收一个 *int 指针
func modifyValueByPtr(ptr *int) {
	if ptr != nil { // 总是一个好习惯去检查指针是否为nil
		*ptr = *ptr * 10 // 修改指针指向的原始值
		fmt.Println("在 modifyValueByPtr内部, *ptr:", *ptr)
	}
}
