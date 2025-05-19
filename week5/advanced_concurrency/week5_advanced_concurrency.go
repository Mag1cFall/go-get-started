package main

import (
	"fmt"
	"sync"
	"time"
)

// --- 1. Channel 深入 ---

// a) 关闭 Channel 的规则回顾和强调
//    - Channel应该由发送方关闭。接收方不应该关闭Channel，因为它们无法知道发送方是否还会发送数据。
//    - 对一个已关闭的Channel进行发送操作会导致panic。
//    - 从一个已关闭的Channel接收数据会立即返回该Channel类型的零值，以及一个表示Channel是否已关闭的布尔值 (v, ok := <-ch)。
//    - 关闭一个nil Channel会导致panic。
//    - 关闭一个已经关闭的Channel会导致panic。

func closeChannelExample(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("  closeChannelExample: ---")
	ch := make(chan int, 3)

	go func() {
		for i := 1; i <= 4; i++ {
			fmt.Printf("    Sender: 发送 %d\n", i)
			ch <- i
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("    Sender: 所有数据发送完毕，关闭channel。")
		close(ch) // 发送方关闭channel
	}()

	time.Sleep(10 * time.Millisecond) // 确保sender先运行一会
	fmt.Println("  Receiver: 准备接收数据...")
	// 使用 for...range 从channel接收数据，循环会在channel关闭后自动结束
	for val := range ch {
		fmt.Printf("  Receiver: 接收到 %d\n", val)
	}
	fmt.Println("  Receiver: Channel 已关闭，for...range 循环结束。")

	// 再次尝试从已关闭的channel接收
	val, ok := <-ch
	fmt.Printf("  Receiver: 再次尝试接收: 值=%d, ok=%t\n", val, ok) // ok 会是 false
}

// b) 单向 Channel (Directional Channels)
//    - chan<- T: 只发送Channel (只能向其发送数据)
//    - <-chan T: 只接收Channel (只能从其接收数据)
//    - 单向Channel用于增强类型安全，明确函数或goroutine对channel的操作权限。

// ping 函数只能发送数据到 pingChan (chan<- string)
// 它会接收一个只接收channel pongChan (<-chan string) 来等待确认
func ping(pingChan chan<- string, pongChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	pingMsg := "ping"
	fmt.Printf("    ping goroutine: 发送 '%s'\n", pingMsg)
	pingChan <- pingMsg
	pongMsg := <-pongChan // 等待pong
	fmt.Printf("    ping goroutine: 收到 '%s'\n", pongMsg)
}

// pong 函数只能从 pingChan (<-chan string) 接收数据
// 它会发送数据到只发送channel pongChan (chan<- string)
func pong(pingChan <-chan string, pongChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	pingMsg := <-pingChan // 等待ping
	fmt.Printf("    pong goroutine: 收到 '%s'\n", pingMsg)
	pongMsg := "pong"
	fmt.Printf("    pong goroutine: 发送 '%s'\n", pongMsg)
	pongChan <- pongMsg
}

func directionalChannelExample(wgOuter *sync.WaitGroup) {
	defer wgOuter.Done()
	fmt.Println("  directionalChannelExample: ---")
	pingChan := make(chan string, 1) // 使用缓冲为1，避免简单示例中的死锁可能
	pongChan := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(2)
	go ping(pingChan, pongChan, &wg)
	go pong(pingChan, pongChan, &wg)
	wg.Wait()
	fmt.Println("  directionalChannelExample: ping-pong 完成。")
}

// c) Channel 死锁场景 (简单示例)
//   - 无缓冲Channel：发送时无接收者，或接收时无发送者。
//   - 主goroutine操作无缓冲channel，但没有其他goroutine配合。
func deadlockExampleSimple() {
	fmt.Println("  deadlockExampleSimple: --- (将导致panic)")
	// ch := make(chan int) // 无缓冲channel
	// ch <- 1 // 死锁！main goroutine发送，但没有其他goroutine在接收
	// fmt.Println(<-ch) // 如果上一行不panic，这一行也会死锁

	// ch := make(chan int)
	// go func() {
	// 	val := <-ch // goroutine尝试接收
	// 	fmt.Println("Goroutine received:", val)
	// }()
	// ch <- 10 // main goroutine发送
	// time.Sleep(100 * time.Millisecond) // 等待goroutine执行
	// fmt.Println("  (以上死锁示例已注释，实际演示时可取消注释一个来观察)")
	fmt.Println("  (死锁示例代码已注释，以避免程序崩溃)")
}

// --- 2. select 语句 ---
// select 语句用于在多个channel操作中进行选择。
// - 它会阻塞，直到其中一个case可以执行。
// - 如果多个case同时就绪，select会随机选择一个执行。
// - default case：如果所有其他case都不能立即执行，则执行default case (实现非阻塞操作)。

func selectExample(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("  selectExample: ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(150 * time.Millisecond)
		ch1 <- "消息来自 ch1"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "消息来自 ch2"
	}()

	// 等待两个消息中的一个
	fmt.Println("  selectExample: 等待 ch1 或 ch2 的消息...")
	for i := 0; i < 2; i++ { // 我们期望收到两条消息
		select {
		case msg1 := <-ch1:
			fmt.Printf("    收到: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("    收到: %s\n", msg2)
		}
	}
	fmt.Println("  selectExample: 两条消息均已收到。")

	// select 与 default (非阻塞)
	fmt.Println("\n  selectExample: 非阻塞 select (default case)")
	nonBlockingCh := make(chan string)
	select {
	case msg := <-nonBlockingCh:
		fmt.Printf("    非阻塞接收: %s\n", msg)
	default:
		fmt.Println("    非阻塞: nonBlockingCh 没有消息。")
	}
	// 尝试非阻塞发送
	// go func() { nonBlockingCh <- "test" }() // 如果有接收者，可以发送
	// time.Sleep(10*time.Millisecond)
	select {
	case nonBlockingCh <- "尝试非阻塞发送":
		fmt.Println("    非阻塞发送成功。")
	default:
		fmt.Println("    非阻塞发送失败 (无缓冲channel，没有接收者)。")
	}

	// select 与超时 (time.After)
	fmt.Println("\n  selectExample: select 与超时")
	timeoutCh := make(chan string, 1) // 使用缓冲channel，否则发送也会阻塞
	go func() {
		time.Sleep(200 * time.Millisecond) // 模拟耗时操作
		timeoutCh <- "操作完成"
	}()

	select {
	case res := <-timeoutCh:
		fmt.Printf("    操作结果: %s\n", res)
	case <-time.After(100 * time.Millisecond): // time.After 返回一个channel，在指定时间后发送当前时间
		fmt.Println("    操作超时! (100ms)")
	}

	// 再试一次，这次操作在超时前完成
	go func() {
		time.Sleep(50 * time.Millisecond)
		timeoutCh <- "快速操作完成"
	}()
	select {
	case res := <-timeoutCh:
		fmt.Printf("    操作结果: %s\n", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("    操作超时! (100ms)")
	}
}

// --- 3. 并发安全与锁 (sync包) ---
// 多个Goroutine并发访问共享资源时，可能会发生竞态条件 (race condition)，导致数据不一致或程序崩溃。
// sync 包提供了同步原语，如 Mutex (互斥锁) 和 RWMutex (读写锁)。

var (
	balance int          // 共享资源
	mutex   sync.Mutex   // 互斥锁，保护 balance
	rwMutex sync.RWMutex // 读写锁，保护 balance (用于读多写少的场景)
	once    sync.Once    // 用于确保某个操作只执行一次
)

func deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()         // 获取互斥锁
	defer mutex.Unlock() // 确保在函数退出时释放锁

	fmt.Printf("  存款 %d, 当前余额 %d", amount, balance)
	balance += amount
	fmt.Printf(" -> 新余额 %d\n", balance)
	time.Sleep(10 * time.Millisecond) // 模拟数据库操作等
}

func readBalance(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()         // 获取读锁 (多个goroutine可以同时持有读锁)
	defer rwMutex.RUnlock() // 释放读锁

	fmt.Printf("  读取余额: %d (使用读写锁)\n", balance)
	time.Sleep(5 * time.Millisecond)
}

func writeBalanceWithRWMutex(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.Lock() // 获取写锁 (独占，其他读写操作都会阻塞)
	defer rwMutex.Unlock()

	fmt.Printf("  写入 %d 到余额 (使用读写锁), 当前余额 %d", amount, balance)
	balance = amount
	fmt.Printf(" -> 新余额 %d\n", balance)
	time.Sleep(10 * time.Millisecond)
}

func initializeConfig() {
	fmt.Println("  (sync.Once) 配置初始化函数被调用了。")
	// 模拟加载配置等只需要执行一次的操作
}

func syncExample(wgOuter *sync.WaitGroup) {
	defer wgOuter.Done()
	fmt.Println("  syncExample: ---")
	balance = 1000 // 初始余额

	// Mutex 示例
	fmt.Println("\n  syncExample: Mutex 示例 (并发存款)")
	var wgDeposit sync.WaitGroup
	for i := 0; i < 5; i++ {
		wgDeposit.Add(1)
		go deposit(100, &wgDeposit)
	}
	wgDeposit.Wait()
	fmt.Printf("  Mutex 示例后最终余额: %d\n", balance)

	// RWMutex 示例
	fmt.Println("\n  syncExample: RWMutex 示例 (并发读写)")
	balance = 500 // 重置余额
	var wgRW sync.WaitGroup
	// 启动多个读操作
	for i := 0; i < 3; i++ {
		wgRW.Add(1)
		go readBalance(&wgRW)
	}
	// 启动一个写操作
	wgRW.Add(1)
	go writeBalanceWithRWMutex(2000, &wgRW)
	// 启动更多读操作
	for i := 0; i < 2; i++ {
		wgRW.Add(1)
		go readBalance(&wgRW)
	}
	wgRW.Wait()
	fmt.Printf("  RWMutex 示例后最终余额: %d\n", balance)

	// sync.Once 示例
	fmt.Println("\n  syncExample: sync.Once 示例")
	var wgOnce sync.WaitGroup
	for i := 0; i < 3; i++ {
		wgOnce.Add(1)
		go func() {
			defer wgOnce.Done()
			once.Do(initializeConfig) // initializeConfig只会被执行一次
			fmt.Println("    Goroutine尝试执行初始化。")
		}()
	}
	wgOnce.Wait()
	// 再次调用 once.Do 不会执行 initializeConfig
	once.Do(initializeConfig)
	fmt.Println("  sync.Once 示例结束。")
}

func main() {
	fmt.Println("--- 第5周学习：并发编程深入 ---")

	var wg sync.WaitGroup // WaitGroup for top-level examples

	fmt.Println("\n--- 1. Channel 深入 ---")
	wg.Add(1)
	go closeChannelExample(&wg)
	wg.Wait()

	wg.Add(1)
	go directionalChannelExample(&wg)
	wg.Wait()

	deadlockExampleSimple() // 内部有注释，不会实际panic

	fmt.Println("\n--- 2. select 语句 ---")
	wg.Add(1)
	go selectExample(&wg)
	wg.Wait()

	fmt.Println("\n--- 3. 并发安全与锁 (sync包) ---")
	wg.Add(1)
	go syncExample(&wg)
	wg.Wait()

	fmt.Println("\n--- 第5周并发编程深入学习结束 ---")
}
