package main

import (
	"fmt"
	"reflect" // 用于反射示例
	// 用于获取 Goroutine 信息等 (可选)
	"sync"
)

// --- 核心原理回顾与深化 ---

// --- 1. make vs new ---
// new(T):
//   - 为类型 T 的新项分配内存。
//   - 将内存初始化为 T 类型的零值。
//   - 返回一个指向该内存地址的指针 (*T)。
// make(T, args):
//   - 仅用于创建和初始化 slice, map, 和 channel 这三种内置的引用类型。
//   - 返回一个已初始化的 (非零值) T 类型的值 (不是指针)。
//   - 对于 slice 和 map，可以指定初始大小/容量。

func makeVsNewExample() {
	fmt.Println("\n--- 1. make vs new ---")

	// new 示例
	var p *int = new(int) // p 是一个 *int 指针, *p 的值是 0 (int的零值)
	fmt.Printf("  new(int): p = %v, *p = %d\n", p, *p)
	*p = 100
	fmt.Printf("  new(int) after assignment: p = %v, *p = %d\n", p, *p)

	type Point struct{ X, Y int }
	var pp *Point = new(Point) // pp 是一个 *Point 指针, *pp 的值是 {0 0} (Point的零值)
	fmt.Printf("  new(Point): pp = %v, *pp = %+v\n", pp, *pp)
	pp.X = 1 // Go 自动解引用: (*pp).X = 1

	// make 示例
	// Slice: make([]T, length, capacity)
	s := make([]int, 3, 5) // s 是一个 []int 切片, 长度3, 容量5, 初始值 [0 0 0]
	fmt.Printf("  make([]int, 3, 5): s = %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Map: make(map[K]V, initialCapacity)
	m := make(map[string]int, 5) // m 是一个 map[string]int, 初始为空但有预分配空间
	fmt.Printf("  make(map[string]int, 5): m = %v, len=%d\n", m, len(m))
	m["a"] = 1

	// Channel: make(chan T, bufferCapacity)
	ch := make(chan int, 1) // ch 是一个缓冲为1的 int 型 channel
	fmt.Printf("  make(chan int, 1): ch = %v\n", ch)
	// ch <- 1 // 可以发送一个值

	// 总结：
	// - new 返回指针，指向零值。
	// - make 返回初始化后的引用类型本身（slice, map, channel）。
}

// --- 2. 函数传结构体：值 vs 指针 ---
type MyStruct struct {
	Value int
	Name  string
}

func modifyStructByValue(s MyStruct, newVal int, newName string) {
	s.Value = newVal // 修改的是副本
	s.Name = newName
	fmt.Printf("    modifyStructByValue (内部): s = %+v\n", s)
}

func modifyStructByPointer(sPtr *MyStruct, newVal int, newName string) {
	if sPtr == nil {
		return
	}
	sPtr.Value = newVal // 修改的是原始结构体
	sPtr.Name = newName
	fmt.Printf("    modifyStructByPointer (内部): sPtr = %+v\n", *sPtr)
}

func structPassExample() {
	fmt.Println("\n--- 2. 函数传结构体：值 vs 指针 ---")
	original := MyStruct{Value: 10, Name: "Original"}
	fmt.Printf("  原始结构体: %+v\n", original)

	modifyStructByValue(original, 20, "ValueCopy")
	fmt.Printf("  值传递后 (原始结构体不变): %+v\n", original)

	modifyStructByPointer(&original, 30, "PointerModified")
	fmt.Printf("  指针传递后 (原始结构体改变): %+v\n", original)

	// 性能：
	// - 值传递：复制整个结构体，对于大结构体开销较大。
	// - 指针传递：只复制指针（通常4或8字节），开销小。
	// 副作用：
	// - 值传递：函数内部修改不影响外部。
	// - 指针传递：函数内部修改会影响外部。
	// 选择：
	// - 如果结构体小且不希望函数修改它，或需要其不可变性，用值传递。
	// - 如果结构体大，或需要在函数内修改它，用指针传递。
	// - Go中，map、slice、channel本身就是引用类型（内部包含指针），通常直接值传递它们即可修改其内容。
}

// --- 3. Goroutine 调度 (PMG模型) ---
// (概念性解释，代码无法直接演示模型本身)
// P: Processor (处理器/上下文)。P 的数量通常等于 CPU 核心数 (GOMAXPROCS)。每个 P 有一个本地可运行 Goroutine 队列 (LRQ)。
// M: Machine (系统线程)。由操作系统管理的线程。M 需要绑定一个 P 才能执行 Go 代码。
// G: Goroutine。Go 的轻量级并发执行单元。包含栈、指令指针等信息。
//
// 协作：
// - M 从 P 的 LRQ 中获取 G 来执行。
// - 如果 P 的 LRQ 为空，P 会尝试从其他 P 的 LRQ "窃取" G，或者从全局可运行队列 (GRQ) 获取 G。
// - 如果 G 发生系统调用 (如文件I/O, 网络I/O) 导致阻塞，M 会与 P 解绑，P 可以被其他 M 使用。
//   当系统调用完成后，原来的 G 会被放回某个 P 的 LRQ 等待再次调度。
// - Go 调度器是用户态的，切换 G 的成本远低于线程切换。
//
// 抢占 (Preemption)：
// - 从 Go 1.14 开始，引入了基于信号的异步抢占机制。
// - 如果一个 G 长时间占用 M 而不主动让出（例如死循环或无阻塞点的计算密集型任务），
//   调度器可以通过发送信号来中断 M，保存当前 G 的状态，并将 G 放回队列，让其他 G 有机会执行。
// - 这使得 Go 的调度更加公平，防止单个 Goroutine 饿死其他 Goroutine。

// --- 4. Goroutine 阻塞情况 ---
// - Channel 操作：
//   - 对无缓冲 channel 发送，若无接收者，阻塞。
//   - 对无缓冲 channel 接收，若无发送者，阻塞。
//   - 对已满的缓冲 channel 发送，阻塞。
//   - 对已空的缓冲 channel 接收，阻塞。
// - select 语句：如果所有 case 都无法立即执行，且没有 default case，则 select 阻塞。
// - sync 包的同步原语：
//   - mutex.Lock()：如果锁已被其他 goroutine 持有，阻塞。
//   - wg.Wait()：如果 WaitGroup 计数器 > 0，阻塞。
// - 网络 I/O：例如等待网络连接、读取网络数据。
// - 文件 I/O：例如读取大文件。
// - time.Sleep()：主动阻塞。
// - 系统调用：某些可能导致当前 M 阻塞的系统调用。

// --- 5. Goroutine 状态 (PMG中) ---
// (概念性，与调度器内部状态相关)
// - _Gidle: 刚被分配，还未初始化。
// - _Grunnable: 在可运行队列中，等待被 M 执行。
// - _Grunning: 正在某个 M 上执行。
// - _Gsyscall: 正在执行一个阻塞的系统调用。
// - _Gwaiting: 因为某种原因（如 channel 操作、锁等待）而阻塞，不在可运行队列中。
// - _Gdead: 已退出或被终止，但其栈空间可能还未被回收。
// (还有其他状态如 _Gcopystack, _Gpreempted 等)

// --- 6. 内存占用：线程 vs Goroutine ---
// - 线程 (OS Thread):
//   - 由操作系统内核调度。
//   - 栈空间通常较大且固定 (例如 Linux 上可能是几 MB，Windows 上默认 1MB)。
//   - 创建和销毁开销较大。
//   - 上下文切换开销较大 (需要内核态参与)。
// - Goroutine:
//   - 由 Go 运行时在用户态调度。
//   - 初始栈空间很小 (例如 Go 1.4+ 是 2KB)，并且可以根据需要动态增长和收缩。
//   - 创建和销毁开销非常小。
//   - 上下文切换开销小 (大部分在用户态完成)。
// 结论：Goroutine 比线程轻量得多，可以轻松创建成千上万甚至数百万个 Goroutine。

// --- 7. Goroutine 资源占用与抢占 (见第3点PMG模型中的抢占) ---

// --- 8. OOM (Out of Memory)：线程OOM vs Goroutine OOM ---
// - 线程 OOM：如果创建过多线程，每个线程都有较大的固定栈空间，很快会耗尽系统内存，导致 OOM。
// - Goroutine OOM：
//   - 虽然 Goroutine 栈小且可增长，但如果创建了极大量的 Goroutine，并且每个 Goroutine 都持有一定资源（如内存、channel引用等），
//     或者 Goroutine 栈因深度递归等原因增长得非常大，也可能耗尽堆内存导致 OOM。
//   - 更常见的是，如果 Goroutine 泄漏（创建后无法正常退出并释放资源），会逐渐消耗内存。

// --- 9. panic 传递：Goroutine 间的 panic ---
// - 一个 Goroutine 中的 panic 如果没有被该 Goroutine 内的 defer + recover 捕获，
//   会导致该 Goroutine 终止，并通常导致整个程序崩溃。
// - panic 不会直接“传递”到另一个 Goroutine。每个 Goroutine 的 panic 是独立的。
// - 如果主 Goroutine panic，整个程序会崩溃。
// - 如果一个子 Goroutine panic 且未被恢复，它会终止，但如果主 Goroutine 或其他 Goroutine 仍在运行，
//   程序可能会继续执行一段时间，但最终通常也会因未处理的 panic 而崩溃（除非 panic 被 recover）。
//   更准确地说，一个未被恢复的 panic 会导致整个程序终止。

func panicPropagationExample() {
	fmt.Println("\n--- 9. panic 传递示例 ---")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("    子 Goroutine 捕获到 panic: %v\n", r)
			}
		}()
		fmt.Println("    子 Goroutine: 准备 panic...")
		panic("子 Goroutine 内部的 panic")
	}()
	wg.Wait()
	fmt.Println("    主 Goroutine: 子 Goroutine 已结束。")

	// 如果子 Goroutine 的 panic 未被捕获，程序会崩溃。
	// go func() {
	// 	fmt.Println("    另一个子 Goroutine: 准备 panic (未捕获)...")
	// 	panic("未捕获的 panic")
	// }()
	// time.Sleep(100 * time.Millisecond) // 给它一点时间 panic
	fmt.Println("    (未捕获 panic 的示例已注释)")
}

// --- 10. defer 与子 Goroutine 的 panic ---
// - 父 Goroutine 中的 defer 语句不能捕获子 Goroutine 中发生的 panic。
// - 每个 Goroutine 都有自己的 defer 栈。panic 只能被其所在 Goroutine 的 defer + recover 捕获。

// --- 11. 反射 (reflect) ---
// reflect 包提供了运行时反射的能力，允许程序在运行时检查变量的类型和值，
// 以及动态调用方法和操作结构体字段。
// 主要类型：
// - reflect.Type: 表示一个 Go 类型。可以通过 reflect.TypeOf(interface{}) 获取。
// - reflect.Value: 表示一个 Go 值。可以通过 reflect.ValueOf(interface{}) 获取。
// 常用方法：
// - Value.Kind(): 获取值的类别 (如 reflect.Int, reflect.String, reflect.Struct, reflect.Ptr 等)。
// - Value.Type(): 获取值的类型 (reflect.Type)。
// - Value.Interface(): 将 reflect.Value 转换回 interface{}。
// - Value.NumField() / Value.Field(i): (用于 Struct) 获取字段数量和特定字段。
// - Value.NumMethod() / Value.Method(i): 获取方法数量和特定方法。
// - Value.Call(args []reflect.Value): 调用方法。
// - Value.Set(newValue reflect.Value): 修改值 (前提是该 Value 是可设置的，通常需要通过指针获取的 Value)。

type ReflectDemo struct {
	Name         string `json:"name_tag" custom:"demo_tag"`
	Age          int    `json:"age_tag"`
	privateField string // 未导出字段不能通过反射直接设置，但可以获取（如果通过非指针的 Value）
}

func (rd ReflectDemo) Greet(prefix string) string {
	return fmt.Sprintf("%s, my name is %s and I am %d.", prefix, rd.Name, rd.Age)
}

func (rd *ReflectDemo) SetAge(newAge int) {
	rd.Age = newAge
}

func reflectionExample() {
	fmt.Println("\n--- 11. 反射 (reflect) ---")
	demo := ReflectDemo{Name: "RooReflect", Age: 5, privateField: "secret"}

	// 获取 Type 和 Value
	t := reflect.TypeOf(demo)
	v := reflect.ValueOf(demo)                                 // 注意：这里 v 是 demo 的副本的 Value
	fmt.Printf("  TypeOf(demo): %v, Kind: %v\n", t, t.Kind())  // main.ReflectDemo, struct
	fmt.Printf("  ValueOf(demo): %v, Kind: %v\n", v, v.Kind()) // {RooReflect 5 secret}, struct

	// 遍历结构体字段
	fmt.Println("  遍历字段 (通过 reflect.Type):")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagJson := field.Tag.Get("json") // 获取 json 标签
		tagCustom := field.Tag.Get("custom")
		fmt.Printf("    字段名: %s, 类型: %v, JSON Tag: '%s', Custom Tag: '%s'\n",
			field.Name, field.Type, tagJson, tagCustom)
	}

	// 获取字段值 (通过 reflect.Value)
	fmt.Println("  获取字段值 (通过 reflect.Value):")
	nameValue := v.FieldByName("Name")
	if nameValue.IsValid() { // 检查字段是否存在
		fmt.Printf("    Name: %s (类型: %s)\n", nameValue.String(), nameValue.Kind())
	}
	// fmt.Println(v.FieldByName("privateField").String()) // 如果直接 .String() 会 panic 因为未导出
	// 但如果用 Interface() 再类型断言，或者知道其类型，可以获取。

	// 修改结构体字段 (需要通过指针的 Value)
	fmt.Println("  修改结构体字段 (需要指针):")
	vp := reflect.ValueOf(&demo) // vp 是指向 demo 的指针的 Value
	// 要修改结构体字段，需要获取元素 (Elem) 的 Value，并且它必须是可设置的 (CanSet)
	if vp.Kind() == reflect.Ptr {
		elem := vp.Elem() // Elem() 获取指针指向的元素
		if elem.Kind() == reflect.Struct {
			nameField := elem.FieldByName("Name")
			if nameField.IsValid() && nameField.CanSet() {
				nameField.SetString("RooReflectModified")
				fmt.Printf("    修改后 Name (通过反射): %s\n", demo.Name)
			} else {
				fmt.Println("    Name 字段不可设置或无效。")
			}
		}
	}

	// 调用方法 (通过 reflect.Value)
	fmt.Println("  调用方法 (通过反射):")
	// Greet 方法是值接收者，可以在 v (值的 Value) 或 vp (指针的 Value) 上调用
	methodGreet := v.MethodByName("Greet")
	if methodGreet.IsValid() {
		args := []reflect.Value{reflect.ValueOf("Hi")}
		result := methodGreet.Call(args) // Call 返回 []reflect.Value
		fmt.Printf("    Greet(\"Hi\") 调用结果: %s\n", result[0].String())
	}

	// SetAge 方法是指针接收者，必须在指针的 Value (vp) 上调用，或者在可寻址的值的 Value 上调用
	methodSetAge := vp.MethodByName("SetAge") // 或者 reflect.ValueOf(&demo).MethodByName("SetAge")
	if methodSetAge.IsValid() {
		args := []reflect.Value{reflect.ValueOf(6)}
		methodSetAge.Call(args)
		fmt.Printf("    SetAge(6) 调用后, demo.Age: %d\n", demo.Age)
	}
	// 反射是强大的工具，但也更复杂，性能开销较大，应谨慎使用。
	// 通常用于编写通用库、编解码、ORM 等场景。
}

// --- 12. 锁机制 (sync.Mutex) ---
// (已在 week5_advanced_concurrency.go 中详细演示)
// - 正常模式 (Non-recursive Mutex): Go 的 Mutex 是非递归的。一个 goroutine 不能重入已持有的锁。
// - 饥饿模式 (Starvation Mode):
//   - 从 Go 1.9 开始，Mutex 引入了饥饿模式以提高公平性。
//   - 当一个 goroutine 等待锁超过1毫秒时，Mutex 进入饥饿模式。
//   - 在饥饿模式下，锁的所有权直接从解锁的 goroutine 传递给等待队列中的第一个 goroutine。
//   - 新到达的 goroutine 不会尝试获取锁，即使锁看起来是可用的（它们不会自旋），而是排在等待队列的尾部。
//   - 如果一个 goroutine 获得了锁并且发现它是等待队列中的最后一个，或者它等待锁的时间少于1毫秒，Mutex 会切换回正常模式。
// - 底层大致实现：
//   - Mutex 内部有一个状态字 (state)，用位操作来表示锁是否被持有、是否有等待者、是否处于饥饿模式等。
//   - 尝试获取锁 (Lock):
//     - 快速路径：通过原子操作尝试设置锁定位。如果成功，立即返回。
//     - 慢速路径：如果锁已被持有，goroutine 会进入等待（可能先自旋一小段时间，然后通过运行时信号量休眠）。
//   - 释放锁 (Unlock):
//     - 通过原子操作清除锁定位。
//     - 如果有等待者，唤醒一个或（饥饿模式下）直接移交。

func main() {
	fmt.Println("--- 第7周学习：核心原理回顾与深化 ---")

	makeVsNewExample()
	structPassExample()

	// PMG, Goroutine状态, 内存占用, OOM, panic传递, defer与子goroutine的panic 等主要通过注释解释
	fmt.Println("\n--- Goroutine 调度、状态、内存、OOM、panic传递等原理见代码注释 ---")
	panicPropagationExample() // 演示panic在goroutine内的捕获

	reflectionExample()

	fmt.Println("\n--- 核心原理回顾学习结束 ---")
	fmt.Println("注意：PMG模型、Goroutine状态、锁的底层实现等是非常深入的主题，这里的解释是高度简化的。")
	fmt.Println("      建议阅读源码或专门的Go运行时分析文章以获得更全面的理解。")
	fmt.Println("      例如 Dave Cheney 的博客, 《Go语言底层原理剖析》等资源。")
}
