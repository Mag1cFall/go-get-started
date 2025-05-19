package main

import (
	"fmt"
	"sync" // 导入 sync 包，用于 WaitGroup
	"time" // 导入 time 包，用于在 goroutine 中模拟耗时操作
)

// --- 1. Goroutine ---
// Goroutine 是 Go 语言中轻量级的并发执行单元。
// 使用 go 关键字后跟一个函数调用，即可启动一个新的 Goroutine。
// main 函数本身也运行在一个 Goroutine 中。

func sayHello() {
	fmt.Println("  sayHello Goroutine: Hello from a new Goroutine!")
}

func printNumbers() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("  printNumbers Goroutine: Number %d\n", i)
		time.Sleep(100 * time.Millisecond) // 模拟一些工作
	}
	fmt.Println("  printNumbers Goroutine: Finished.")
}

func printLetters() {
	for charCode := 'a'; charCode <= 'c'; charCode++ {
		fmt.Printf("  printLetters Goroutine: Letter %c\n", charCode)
		time.Sleep(150 * time.Millisecond) // 模拟一些工作
	}
	fmt.Println("  printLetters Goroutine: Finished.")
}

func main() {
	fmt.Println("--- 第4周学习：并发编程初步 (Goroutines, Channels, WaitGroup) ---")

	// --- Goroutine 示例 ---
	fmt.Println("\n--- 1. Goroutine 示例 ---")
	go sayHello() // 启动一个新的 Goroutine 来执行 sayHello

	// 注意：当 main Goroutine 结束时，程序会立即退出，
	// 即使其他 Goroutine 可能还没有执行完毕。
	// 为了观察 sayHello 的输出，我们可能需要让 main Goroutine 等待一会儿。
	// （更好的方式是使用 sync.WaitGroup 或 Channel，后面会讲）
	fmt.Println("main Goroutine: sayHello Goroutine launched.")
	time.Sleep(50 * time.Millisecond) // 等待一下，让 sayHello 有机会执行

	fmt.Println("\n启动多个 Goroutines:")
	go printNumbers() // 启动 printNumbers Goroutine
	go printLetters() // 启动 printLetters Goroutine

	fmt.Println("main Goroutine: printNumbers 和 printLetters Goroutines launched.")
	// 同样，需要等待，否则 main 可能先结束
	// 输出的顺序可能是不确定的，因为 Goroutine 是并发执行的。
	time.Sleep(1 * time.Second) // 等待足够长的时间让它们完成

	// --- 2. sync.WaitGroup ---
	// sync.WaitGroup 用于等待一组 Goroutine 完成。
	// - Add(delta int): 增加计数器，表示有多少个 Goroutine 需要等待。
	// - Done(): 减少计数器，通常在 Goroutine 完成时通过 defer 调用。
	// - Wait(): 阻塞，直到计数器变为零。
	fmt.Println("\n--- 2. 使用 sync.WaitGroup 等待 Goroutines ---")
	var wg sync.WaitGroup // 创建一个 WaitGroup

	numTasks := 3
	wg.Add(numTasks) // 设置需要等待的 Goroutine 数量

	for i := 1; i <= numTasks; i++ {
		// 启动 Goroutine
		go func(taskID int) {
			defer wg.Done() // 任务完成时，调用 Done() 使计数器减1
			fmt.Printf("  WaitGroup Task %d: Starting...\n", taskID)
			time.Sleep(time.Duration(taskID*100) * time.Millisecond) // 模拟不同耗时的任务
			fmt.Printf("  WaitGroup Task %d: Finished.\n", taskID)
		}(i) // 将 i 作为参数传递给匿名函数，避免闭包问题
	}

	fmt.Println("main Goroutine: 所有 WaitGroup tasks 已启动，等待完成...")
	wg.Wait() // 等待所有 Goroutine 调用 Done()，即计数器归零
	fmt.Println("main Goroutine: 所有 WaitGroup tasks 已完成!")

	// --- 3. Channel (通道) ---
	// Channel 是类型化的管道，用于在 Goroutine 之间传递数据，从而实现通信和同步。
	// 操作符 <- 用于在 Channel 上发送或接收数据。
	// ch <- v    // 发送 v 到 Channel ch.
	// v := <-ch  // 从 Channel ch 接收数据并赋值给 v.
	// (数据流向箭头的方向)

	fmt.Println("\n--- 3. Channel 示例 ---")

	// a) 无缓冲 Channel (Unbuffered Channel)
	//    - 发送操作会阻塞，直到另一个 Goroutine 在同一 Channel 上进行接收操作。
	//    - 接收操作会阻塞，直到另一个 Goroutine 在同一 Channel 上进行发送操作。
	//    - 用于 Goroutine 之间的同步。
	fmt.Println("  --- a) 无缓冲 Channel ---")
	messageChannel := make(chan string) // 创建一个无缓冲的 string 类型 Channel

	go func() {
		fmt.Println("    Sender Goroutine: 准备发送消息...")
		time.Sleep(200 * time.Millisecond)
		messageChannel <- "Hello from无缓冲Channel!" // 发送消息到 Channel
		fmt.Println("    Sender Goroutine: 消息已发送。")
	}()

	fmt.Println("  main Goroutine: 等待从 Channel 接收消息...")
	receivedMessage := <-messageChannel // 从 Channel 接收消息 (会阻塞)
	fmt.Printf("  main Goroutine: 接收到消息: \"%s\"\n", receivedMessage)
	// close(messageChannel) // Channel 使用完毕后可以关闭，但并非总是必须

	// b) 有缓冲 Channel (Buffered Channel)
	//    - make(chan Type, capacity)
	//    - 发送操作仅在缓冲区满时阻塞。
	//    - 接收操作仅在缓冲区空时阻塞。
	fmt.Println("\n  --- b) 有缓冲 Channel ---")
	bufferedChan := make(chan int, 2) // 创建一个容量为2的 int 类型有缓冲 Channel

	go func() {
		fmt.Println("    Buffered Sender: 发送 1...")
		bufferedChan <- 1
		fmt.Println("    Buffered Sender: 发送 1 完成。")
		fmt.Println("    Buffered Sender: 发送 2...")
		bufferedChan <- 2
		fmt.Println("    Buffered Sender: 发送 2 完成。")
		fmt.Println("    Buffered Sender: 尝试发送 3 (缓冲区已满，会阻塞)...")
		bufferedChan <- 3 // 这会阻塞，直到有接收者取走数据
		fmt.Println("    Buffered Sender: 发送 3 完成。")
		close(bufferedChan) // 当所有数据都发送完毕后，发送方可以关闭 Channel
		// 关闭 Channel 表示不会再有新的值发送到这个 Channel。
		// 接收方仍然可以从已关闭的 Channel 中读取已发送的值。
		fmt.Println("    Buffered Sender: Channel 已关闭。")
	}()

	time.Sleep(100 * time.Millisecond) // 给发送方一点时间先填满缓冲区

	fmt.Println("  main Goroutine: 从有缓冲 Channel 接收数据...")
	fmt.Printf("  接收到: %d\n", <-bufferedChan)
	fmt.Printf("  接收到: %d\n", <-bufferedChan)
	// 此时发送方的 bufferedChan <- 3 应该可以成功了
	time.Sleep(100 * time.Millisecond) // 等待发送方发送第3个并关闭

	// 从已关闭的 Channel 接收数据
	// 可以使用 for...range 循环来接收 Channel 中的所有值，直到 Channel 关闭。
	fmt.Println("  main Goroutine: 使用 for...range 从 (可能已关闭的) Channel 接收剩余数据:")
	// 在这个例子中，因为我们知道发送方会发送第三个值然后关闭，
	// 所以直接再接收一次，或者用 for range
	val, ok := <-bufferedChan // 检查 Channel 是否已关闭
	if ok {
		fmt.Printf("  再次接收到: %d (ok=%t)\n", val, ok)
	} else {
		fmt.Printf("  Channel 已关闭，无法再接收新值 (ok=%t)\n", ok)
	}
	// 如果 Channel 已经被关闭，并且缓冲区为空，则接收操作会立即返回一个零值和 false。
	val, ok = <-bufferedChan
	fmt.Printf("  尝试从已关闭且已空的 Channel 再次接收: 值=%d, ok=%t\n", val, ok)

	fmt.Println("\n--- 并发编程初步学习结束 ---")
}
