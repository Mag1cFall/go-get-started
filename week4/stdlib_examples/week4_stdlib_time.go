package main

import (
	"fmt"
	"time" // 导入 time 包
)

func main() {
	fmt.Println("--- 第4周学习：常用标准库 (time 包) ---")

	// --- 1. 获取当前时间 ---
	fmt.Println("\n--- 1. 获取当前时间 ---")
	now := time.Now() // time.Time 类型
	fmt.Println("当前时间 (time.Now()):", now)

	// --- 2. 时间的组成部分 ---
	fmt.Println("\n--- 2. 时间的组成部分 ---")
	fmt.Println("  年:", now.Year())
	fmt.Println("  月 (Month类型):", now.Month()) // time.Month 类型 (e.g., January)
	fmt.Println("  月 (数字):", int(now.Month())) // 转换为 int
	fmt.Println("  日:", now.Day())
	fmt.Println("  时:", now.Hour())
	fmt.Println("  分:", now.Minute())
	fmt.Println("  秒:", now.Second())
	fmt.Println("  纳秒:", now.Nanosecond())
	fmt.Println("  星期几 (Weekday类型):", now.Weekday()) // time.Weekday 类型 (e.g., Sunday)
	fmt.Println("  时区:", now.Location())             // *time.Location 类型

	// --- 3. 格式化时间 (Format) ---
	// Go 使用一种特殊的参考时间来定义格式： Mon Jan 2 15:04:05 MST 2006 (即 1月2日 下午3点4分5秒 2006年 MST时区)
	// 你需要按照这个参考时间的格式来指定你想要的输出格式。
	fmt.Println("\n--- 3. 格式化时间 ---")
	fmt.Println("默认格式 (now.String()):", now.String())
	// 自定义格式
	fmt.Println("自定义格式 (YYYY-MM-DD HH:MM:SS):", now.Format("2006-01-02 15:04:05"))
	fmt.Println("自定义格式 (YYYY/MM/DD):", now.Format("2006/01/02"))
	fmt.Println("自定义格式 (HH:MM):", now.Format("15:04"))
	fmt.Println("带时区的格式:", now.Format("2006-01-02 15:04:05 MST"))
	// 一些预定义的格式常量
	fmt.Println("RFC3339 格式:", now.Format(time.RFC3339))
	fmt.Println("Kitchen 格式 (小时:分钟 AM/PM):", now.Format(time.Kitchen))

	// --- 4. 解析时间字符串 (Parse) ---
	// Parse(layout, value string) (Time, error)
	// layout 必须是定义格式时使用的参考时间格式。
	fmt.Println("\n--- 4. 解析时间字符串 ---")
	timeStr1 := "2023-10-26 10:30:00"
	layout1 := "2006-01-02 15:04:05"
	parsedTime1, err := time.Parse(layout1, timeStr1)
	if err != nil {
		fmt.Printf("解析时间字符串 '%s' 错误: %v\n", timeStr1, err)
	} else {
		fmt.Printf("解析 '%s' 得到的时间: %v\n", timeStr1, parsedTime1)
	}

	timeStr2 := "26/Oct/2023 08:15PM"
	layout2 := "02/Jan/2006 03:04PM" // 注意 PM
	parsedTime2, err := time.Parse(layout2, timeStr2)
	if err != nil {
		fmt.Printf("解析时间字符串 '%s' 错误: %v\n", timeStr2, err)
	} else {
		fmt.Printf("解析 '%s' 得到的时间: %v\n", timeStr2, parsedTime2)
	}

	// ParseInLocation: 可以在指定时区解析时间
	timeStrLocal := "2023-10-26 14:00:00"
	loc, _ := time.LoadLocation("America/New_York") // 加载时区
	parsedTimeInLoc, err := time.ParseInLocation(layout1, timeStrLocal, loc)
	if err != nil {
		fmt.Printf("在指定时区解析 '%s' 错误: %v\n", timeStrLocal, err)
	} else {
		fmt.Printf("在纽约时区解析 '%s' 得到的时间: %v\n", timeStrLocal, parsedTimeInLoc)
	}

	// --- 5. 时间点操作 ---
	fmt.Println("\n--- 5. 时间点操作 ---")
	// 创建特定时间点
	specificTime := time.Date(2025, time.January, 1, 12, 0, 0, 0, time.UTC) // 年,月,日,时,分,秒,纳秒,时区
	fmt.Println("特定时间点 (UTC):", specificTime)

	// 时间点比较
	fmt.Println("  now.Before(specificTime):", now.Before(specificTime))
	fmt.Println("  now.After(specificTime):", now.After(specificTime))
	fmt.Println("  now.Equal(specificTime):", now.Equal(specificTime))

	// 时间点加减 (Duration)
	oneHour := time.Hour // time.Duration 类型，预定义了 Hour, Minute, Second, Millisecond, Microsecond, Nanosecond
	oneDay := 24 * time.Hour
	tomorrow := now.Add(oneDay)
	yesterday := now.Add(-oneDay) // 或 now.AddDate(0,0,-1)
	nextHour := now.Add(oneHour)
	fmt.Println("  当前时间:", now.Format(time.RFC1123))
	fmt.Println("  一小时后:", nextHour.Format(time.RFC1123))
	fmt.Println("  明天:", tomorrow.Format(time.RFC1123))
	fmt.Println("  昨天:", yesterday.Format(time.RFC1123))

	// AddDate(years, months, days int)
	nextMonth := now.AddDate(0, 1, 0)
	fmt.Println("  下个月的今天:", nextMonth.Format("2006-01-02"))

	// Sub(t2 Time) Duration: 计算两个时间点之间的差值
	diff := specificTime.Sub(now)
	fmt.Printf("  specificTime 和 now 之间相差: %v (约 %.2f 小时)\n", diff, diff.Hours())

	// --- 6. 时间戳 (Unix Timestamp) ---
	// Unix 时间戳是指自 1970年1月1日（UTC）起经过的秒数或纳秒数。
	fmt.Println("\n--- 6. 时间戳 ---")
	fmt.Println("当前时间的 Unix 秒数:", now.Unix())
	fmt.Println("当前时间的 Unix 纳秒数:", now.UnixNano())

	// 从 Unix 时间戳创建 time.Time
	unixTimestamp := int64(1672531200)                         // 2023-01-01 00:00:00 UTC
	timeFromUnix := time.Unix(unixTimestamp, 0)                // 第二个参数是纳秒部分
	fmt.Println("从 Unix 时间戳还原的时间:", timeFromUnix.In(time.UTC)) // 显示为 UTC

	// --- 7. 睡眠 (Sleep) ---
	fmt.Println("\n--- 7. 睡眠 (Sleep) ---")
	fmt.Println("准备睡眠 1 秒钟...")
	// time.Sleep(1 * time.Second) // 会使程序暂停执行
	fmt.Println("  (睡眠操作已注释，以避免执行流程暂停)")
	fmt.Println("睡眠结束 (如果未注释)。")

	// --- 8. 定时器 (Timer) 和 打点器 (Ticker) ---
	// Timer: 在指定时间后触发一次事件。
	// Ticker: 按固定的时间间隔重复触发事件。
	// (这些更常用于并发编程，这里仅作简单提及)
	fmt.Println("\n--- 8. 定时器和打点器 (初步提及) ---")
	// timer := time.NewTimer(2 * time.Second)
	// <-timer.C // 阻塞直到定时器触发
	// fmt.Println("  2秒定时器触发了! (如果未注释)")

	// ticker := time.NewTicker(1 * time.Second)
	// go func() {
	// 	for t := range ticker.C {
	// 		fmt.Println("  打点器在", t, "触发 (如果未注释)")
	// 	}
	// }()
	// time.Sleep(3 * time.Second) // 让打点器运行一会儿
	// ticker.Stop()
	// fmt.Println("  打点器已停止 (如果未注释)")
	fmt.Println("  (定时器和打点器相关代码已注释，它们通常用于并发场景)")

	fmt.Println("\n--- time 包学习结束 ---")
}
