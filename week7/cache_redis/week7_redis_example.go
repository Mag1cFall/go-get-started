package main

import (
	"context" // go-redis/redis v8+ 需要 context.Context
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 导入 go-redis 客户端库
)

// Redis 连接选项
// 通常 Redis 默认运行在 localhost:6379，没有密码。
// 如果你的 Redis 配置不同，请修改这些值。
const (
	redisAddr     = "localhost:6379"
	redisPassword = "" // 如果你的 Redis 没有密码，则留空
	redisDB       = 0  // 默认数据库
)

func main() {
	fmt.Println("--- 第7周学习：缓存操作 (Redis 与 go-redis/redis) ---")
	fmt.Println("!!! 重要: 请确保你已启动 Redis 服务器，并根据需要更新了代码中的连接常量。")

	// --- 1. 创建 Redis 客户端 ---
	// go-redis/redis v8 使用 redis.NewClient 创建客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

	// 使用 context.Context (通常是 background context 或带有超时的 context)
	// 对于简单的示例，Background context 即可。
	ctx := context.Background()

	// --- 2. 测试连接 (Ping) ---
	// Ping 命令用于检查与 Redis 服务器的连接是否正常。
	pong, err := rdb.Ping(ctx).Result() // .Result() 会阻塞并返回命令的结果和错误
	if err != nil {
		fmt.Printf("  [Redis] 连接 Redis 失败 (rdb.Ping): %v\n", err)
		fmt.Println("  请确保 Redis 服务器正在运行，并且地址/密码配置正确。")
		return // 如果无法连接，后续操作无意义
	}
	fmt.Printf("  [Redis] 成功连接到 Redis! Ping 响应: %s\n", pong)

	// --- 3. 基本的 SET 和 GET 操作 (String 类型) ---
	fmt.Println("\n--- 3. String 类型操作 (SET, GET) ---")
	key := "myKey"
	value := "Hello Redis from Go!"
	expiration := 10 * time.Minute // 设置键的过期时间 (可选，0表示不过期)

	// SET key value [EX seconds | PX milliseconds | KEEPTTL]
	err = rdb.Set(ctx, key, value, expiration).Err() // .Err() 返回命令执行的错误
	if err != nil {
		fmt.Printf("  [Redis] SET key '%s' 失败: %v\n", key, err)
	} else {
		fmt.Printf("  [Redis] 成功 SET key '%s' = '%s' (过期时间: %v)\n", key, value, expiration)
	}

	// GET key
	retrievedValue, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil { // redis.Nil 表示键不存在
		fmt.Printf("  [Redis] GET key '%s': 键不存在\n", key)
	} else if err != nil {
		fmt.Printf("  [Redis] GET key '%s' 失败: %v\n", key, err)
	} else {
		fmt.Printf("  [Redis] 成功 GET key '%s' = '%s'\n", key, retrievedValue)
	}

	// GET 一个不存在的键
	nonExistentKey := "doesNotExistKey"
	retrievedNonExistent, err := rdb.Get(ctx, nonExistentKey).Result()
	if err == redis.Nil {
		fmt.Printf("  [Redis] GET key '%s': 键不存在 (符合预期)\n", nonExistentKey)
	} else if err != nil {
		fmt.Printf("  [Redis] GET key '%s' 失败: %v\n", nonExistentKey, err)
	} else {
		fmt.Printf("  [Redis] GET key '%s' = '%s' (不应到这里)\n", nonExistentKey, retrievedNonExistent)
	}

	// --- 4. 其他常用操作 (简单提及) ---
	fmt.Println("\n--- 4. 其他常用操作 (简单提及) ---")

	// DEL key [key ...] - 删除键
	keysToDelete := []string{key, "anotherKey"}
	numDeleted, err := rdb.Del(ctx, keysToDelete...).Result() // ... 用于将切片展开为多个参数
	if err != nil {
		fmt.Printf("  [Redis] DEL keys %v 失败: %v\n", keysToDelete, err)
	} else {
		fmt.Printf("  [Redis] 成功 DEL %d 个键: %v\n", numDeleted, keysToDelete)
	}

	// EXISTS key [key ...] - 检查键是否存在
	exists, err := rdb.Exists(ctx, key).Result() // key 已经被删了
	if err != nil {
		fmt.Printf("  [Redis] EXISTS key '%s' 失败: %v\n", key, err)
	} else {
		fmt.Printf("  [Redis] key '%s' 是否存在? %t (0表示不存在, 1表示存在)\n", key, exists == 1)
	}

	// Redis 支持多种数据类型，如 Hash, List, Set, Sorted Set (ZSet)
	// go-redis/redis 为每种类型都提供了相应的操作命令。
	// 例如:
	// - HSet, HGet (Hash)
	// - LPush, RPush, LRange (List)
	// - SAdd, SMembers (Set)
	// - ZAdd, ZRange (Sorted Set)
	fmt.Println("  [Redis] go-redis/redis 支持 Hash, List, Set, Sorted Set 等多种数据类型。")
	fmt.Println("  例如: rdb.HSet(ctx, \"myhash\", \"field1\", \"value1\")")
	fmt.Println("        rdb.LPush(ctx, \"mylist\", \"item1\", \"item2\")")

	// --- 5. 关闭 Redis 客户端连接 (可选) ---
	// 通常，Redis 客户端实例可以被长期持有并在应用程序的生命周期内复用。
	// 如果确实需要关闭，可以调用 Close()。
	// err = rdb.Close()
	// if err != nil {
	// 	fmt.Printf("  [Redis] 关闭 Redis 客户端错误: %v\n", err)
	// } else {
	// 	fmt.Println("  [Redis] Redis 客户端已关闭。")
	// }
	fmt.Println("  (Redis 客户端通常可复用，此处未显式关闭)")

	fmt.Println("\n--- Redis 缓存操作学习结束 ---")
}
