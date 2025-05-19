package main

import (
	"encoding/json" // 导入 json 包
	"fmt"
	"os" // 用于更复杂的示例，如读写JSON文件
)

// --- 1. 定义用于 JSON 操作的结构体 ---
// 结构体标签 (struct tags) 用于控制 JSON 序列化和反序列化的行为。
// `json:"fieldName"`: 指定 JSON 中的字段名。
// `json:"fieldName,omitempty"`: 如果字段值为其类型的零值，则在序列化时忽略该字段。
// `json:"-"`: 在序列化和反序列化时始终忽略该字段。

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email,omitempty"` // 如果 Email 为空字符串，则忽略
	Password string   `json:"-"`               // 密码字段不应包含在JSON中
	IsActive bool     `json:"isActive"`        // Go的bool会转为JSON的true/false
	Profile  Profile  `json:"profileInfo"`     // 嵌套结构体
	Tags     []string `json:"tags,omitempty"`  // 切片
}

type Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarURL string `json:"avatarUrl,omitempty"`
}

func main() {
	fmt.Println("--- 第4周学习：常用标准库 (encoding/json 包) ---")

	// --- 2. 序列化 (Marshalling): Go 结构体 -> JSON 字符串 ---
	fmt.Println("\n--- 2. 序列化 (Go -> JSON) ---")
	user1 := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john.doe@example.com", // Email 非空，会包含
		Password: "supersecret",          // Password 会被忽略
		IsActive: true,
		Profile: Profile{
			FirstName: "John",
			LastName:  "Doe",
			AvatarURL: "http://example.com/avatar.png",
		},
		Tags: []string{"go", "developer", "json"},
	}

	user1JSON, err := json.Marshal(user1) // 返回字节切片 []byte 和 error
	if err != nil {
		fmt.Printf("  json.Marshal (user1) 错误: %v\n", err)
	} else {
		fmt.Printf("  User1 序列化为 JSON: %s\n", string(user1JSON))
	}

	// 使用 MarshalIndent 进行格式化 (带缩进) 的 JSON 输出
	user1JSONFormatted, err := json.MarshalIndent(user1, "", "  ") // prefix="", indent="  " (两个空格)
	if err != nil {
		fmt.Printf("  json.MarshalIndent (user1) 错误: %v\n", err)
	} else {
		fmt.Printf("  User1 格式化 JSON:\n%s\n", string(user1JSONFormatted))
	}

	user2 := User{
		ID:       2,
		Username: "jane_doe",
		// Email 为空，Password 被忽略，AvatarURL 为空，Tags 为 nil
		IsActive: false,
		Profile: Profile{
			FirstName: "Jane",
			LastName:  "Doe",
		},
	}
	user2JSON, _ := json.MarshalIndent(user2, "", "  ")
	fmt.Printf("  User2 (有 omitempty 字段为空) 格式化 JSON:\n%s\n", string(user2JSON))

	// --- 3. 反序列化 (Unmarshalling): JSON 字符串 -> Go 结构体 ---
	fmt.Println("\n--- 3. 反序列化 (JSON -> Go) ---")
	jsonStr := `{"id":101,"username":"test_user","email":"test@example.com","isActive":true,"profileInfo":{"firstName":"Test","lastName":"User"},"tags":["test","sample"]}`

	var decodedUser User
	// json.Unmarshal 需要一个字节切片和一个指向目标结构体的指针
	err = json.Unmarshal([]byte(jsonStr), &decodedUser)
	if err != nil {
		fmt.Printf("  json.Unmarshal 错误: %v\n", err)
	} else {
		fmt.Printf("  从 JSON 反序列化的 User: %+v\n", decodedUser) // %+v 打印字段名和值
		fmt.Printf("    Username: %s, Email: %s\n", decodedUser.Username, decodedUser.Email)
		fmt.Printf("    Profile FirstName: %s\n", decodedUser.Profile.FirstName)
		fmt.Printf("    Tags: %v\n", decodedUser.Tags)
		// Password 字段因为有 `json:"-"` 标签，所以不会被填充
		fmt.Printf("    Password (应为空): '%s'\n", decodedUser.Password)
	}

	// --- 4. 处理任意/未知结构的 JSON (map[string]interface{}) ---
	// 当 JSON 结构不固定或预先未知时，可以将其反序列化到 map[string]interface{}
	fmt.Println("\n--- 4. 处理任意结构的 JSON ---")
	arbitraryJSONStr := `{"name":"Widget","price":19.99,"available":true,"dimensions":{"height":5,"width":10},"colors":["red","blue"]}`

	var arbitraryData map[string]interface{} // interface{} 可以是任何类型
	err = json.Unmarshal([]byte(arbitraryJSONStr), &arbitraryData)
	if err != nil {
		fmt.Printf("  反序列化任意 JSON 错误: %v\n", err)
	} else {
		fmt.Println("  反序列化的任意 JSON 数据:")
		for key, value := range arbitraryData {
			fmt.Printf("    键: %s, 值: %v (类型: %T)\n", key, value, value)
		}
		// 访问特定字段需要类型断言
		if name, ok := arbitraryData["name"].(string); ok {
			fmt.Printf("    提取的 name: %s\n", name)
		}
		if dimensions, ok := arbitraryData["dimensions"].(map[string]interface{}); ok {
			if height, ok := dimensions["height"].(float64); ok { // JSON 数字默认解析为 float64
				fmt.Printf("    提取的 dimensions.height: %.0f\n", height)
			}
		}
	}

	// --- 5. JSON 数组的序列化和反序列化 ---
	fmt.Println("\n--- 5. JSON 数组 ---")
	users := []User{user1, user2}
	usersJSON, _ := json.MarshalIndent(users, "", "  ")
	fmt.Printf("  Users 数组序列化为 JSON:\n%s\n", string(usersJSON))

	jsonArrayStr := `[{"id":201,"username":"userA"},{"id":202,"username":"userB","isActive":true}]`
	var decodedUsers []User
	err = json.Unmarshal([]byte(jsonArrayStr), &decodedUsers)
	if err != nil {
		fmt.Printf("  反序列化 JSON 数组错误: %v\n", err)
	} else {
		fmt.Println("  从 JSON 数组反序列化的 Users:")
		for i, u := range decodedUsers {
			fmt.Printf("    User %d: %+v\n", i+1, u)
		}
	}

	// --- 6. 使用 Encoder 和 Decoder (用于流式处理) ---
	// json.NewEncoder(io.Writer) 和 json.NewDecoder(io.Reader)
	// 适用于处理网络连接、文件等 io.Reader/Writer 流。
	// 这里简单演示写入到 os.Stdout (标准输出)
	fmt.Println("\n--- 6. Encoder / Decoder (简单演示) ---")
	fmt.Println("  使用 Encoder 将 user1 写入 os.Stdout:")
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ") // 设置缩进以便美观输出
	err = encoder.Encode(user1) // Encode 会自动在末尾添加换行符
	if err != nil {
		fmt.Printf("  Encoder.Encode 错误: %v\n", err)
	}

	// Decoder 示例需要一个 io.Reader，例如 strings.NewReader 或 os.File
	// (此处略过 Decoder 的完整示例以保持简洁，其用法与 Unmarshal 类似但针对流)

	fmt.Println("\n--- encoding/json 包学习结束 ---")
}
