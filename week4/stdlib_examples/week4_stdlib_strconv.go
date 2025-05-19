package main

import (
	"fmt"
	"strconv" // 导入 strconv 包
)

func main() {
	fmt.Println("--- 第4周学习：常用标准库 (strconv 包) ---")

	// --- 1. 字符串转换为数值类型 ---
	fmt.Println("\n--- 1. 字符串转换为数值类型 ---")

	// Atoi (ASCII to Integer): string to int
	// 等价于 ParseInt(s, 10, 0)
	sInt := "12345"
	numInt, err := strconv.Atoi(sInt)
	if err != nil {
		fmt.Printf("strconv.Atoi(\"%s\") 错误: %v\n", sInt, err)
	} else {
		fmt.Printf("strconv.Atoi(\"%s\") = %d (类型: %T)\n", sInt, numInt, numInt)
	}

	// ParseInt(s string, base int, bitSize int) (int64, error)
	// base: 2 至 36。如果 base 为 0，则会根据字符串前缀推断进制 ("0x" 为16进制, "0" 为8进制, 否则为10进制)。
	// bitSize: 指定结果必须能容纳的整数类型大小 (0, 8, 16, 32, 64 分别对应 int, int8, int16, int32, int64)。
	sHex := "FF"                                  // 十六进制
	numHex, err := strconv.ParseInt(sHex, 16, 64) // 按16进制解析，结果为 int64
	if err != nil {
		fmt.Printf("strconv.ParseInt(\"%s\", 16, 64) 错误: %v\n", sHex, err)
	} else {
		fmt.Printf("strconv.ParseInt(\"%s\", 16, 64) = %d (十进制表示)\n", sHex, numHex)
	}

	sBinary := "10101"                                 // 二进制
	numBinary, err := strconv.ParseInt(sBinary, 2, 32) // 按2进制解析，结果为 int32
	if err != nil {
		fmt.Printf("strconv.ParseInt(\"%s\", 2, 32) 错误: %v\n", sBinary, err)
	} else {
		fmt.Printf("strconv.ParseInt(\"%s\", 2, 32) = %d\n", sBinary, numBinary)
	}

	// ParseUint(s string, base int, bitSize int) (uint64, error) - 类似 ParseInt，但用于无符号整数
	sUint := "255"
	numUint, err := strconv.ParseUint(sUint, 10, 8) // 结果为 uint8
	if err != nil {
		fmt.Printf("strconv.ParseUint(\"%s\", 10, 8) 错误: %v\n", sUint, err)
	} else {
		fmt.Printf("strconv.ParseUint(\"%s\", 10, 8) = %d\n", sUint, numUint)
	}

	// ParseFloat(s string, bitSize int) (float64, error)
	// bitSize: 32 (float32) 或 64 (float64)
	sFloat := "3.1415926535"
	numFloat64, err := strconv.ParseFloat(sFloat, 64)
	if err != nil {
		fmt.Printf("strconv.ParseFloat(\"%s\", 64) 错误: %v\n", sFloat, err)
	} else {
		fmt.Printf("strconv.ParseFloat(\"%s\", 64) = %f (类型: %T)\n", sFloat, numFloat64, numFloat64)
	}

	sFloat32 := "2.718"
	numFloat32, err := strconv.ParseFloat(sFloat32, 32) // 解析为 float32 精度，但返回 float64
	if err != nil {
		fmt.Printf("strconv.ParseFloat(\"%s\", 32) 错误: %v\n", sFloat32, err)
	} else {
		fmt.Printf("strconv.ParseFloat(\"%s\", 32) = %f (实际存储为 float64, 但按 float32 精度解析)\n", sFloat32, numFloat32)
	}

	// --- 2. 字符串转换为布尔类型 ---
	fmt.Println("\n--- 2. 字符串转换为布尔类型 ---")
	// ParseBool(str string) (bool, error)
	// 接受 "1", "t", "T", "true", "TRUE", "True" 作为 true
	// 接受 "0", "f", "F", "false", "FALSE", "False" 作为 false
	sBoolTrue := "true"
	valBoolTrue, err := strconv.ParseBool(sBoolTrue)
	if err != nil {
		fmt.Printf("strconv.ParseBool(\"%s\") 错误: %v\n", sBoolTrue, err)
	} else {
		fmt.Printf("strconv.ParseBool(\"%s\") = %t\n", sBoolTrue, valBoolTrue)
	}

	sBoolFalse := "F"
	valBoolFalse, err := strconv.ParseBool(sBoolFalse)
	if err != nil {
		fmt.Printf("strconv.ParseBool(\"%s\") 错误: %v\n", sBoolFalse, err)
	} else {
		fmt.Printf("strconv.ParseBool(\"%s\") = %t\n", sBoolFalse, valBoolFalse)
	}

	// --- 3. 数值类型转换为字符串 ---
	fmt.Println("\n--- 3. 数值类型转换为字符串 ---")
	// Itoa (Integer to ASCII): int to string
	// 等价于 FormatInt(int64(i), 10)
	intVal := -456
	strVal := strconv.Itoa(intVal)
	fmt.Printf("strconv.Itoa(%d) = \"%s\" (类型: %T)\n", intVal, strVal, strVal)

	// FormatInt(i int64, base int) string
	// base: 2 至 36
	var int64Val int64 = 255
	fmt.Printf("strconv.FormatInt(%d, 16) (十六进制) = \"%s\"\n", int64Val, strconv.FormatInt(int64Val, 16)) // ff
	fmt.Printf("strconv.FormatInt(%d, 2)  (二进制)   = \"%s\"\n", int64Val, strconv.FormatInt(int64Val, 2)) // 11111111

	// FormatUint(i uint64, base int) string - 类似 FormatInt，但用于无符号整数
	var uint64Val uint64 = 255
	fmt.Printf("strconv.FormatUint(%d, 16) = \"%s\"\n", uint64Val, strconv.FormatUint(uint64Val, 16))

	// FormatFloat(f float64, fmt byte, prec int, bitSize int) string
	// fmt: 格式化方式 ('b' 二进制指数, 'e'/'E' 科学计数法, 'f' 小数点形式, 'g'/'G' 自动选择e或f, 'x' 十六进制指数)
	// prec: 精度 (对 'e', 'f', 'g' 是小数点后的位数；对 'b', 'x' 是有效数字位数)
	// bitSize: 32 (float32) 或 64 (float64)
	floatVal := 3.1415926535
	fmt.Printf("strconv.FormatFloat(%.10f, 'f', 4, 64) (保留4位小数) = \"%s\"\n", floatVal, strconv.FormatFloat(floatVal, 'f', 4, 64))
	fmt.Printf("strconv.FormatFloat(%.10f, 'e', 5, 64) (科学计数法,5位小数) = \"%s\"\n", floatVal, strconv.FormatFloat(floatVal, 'e', 5, 64))

	// --- 4. 布尔类型转换为字符串 ---
	fmt.Println("\n--- 4. 布尔类型转换为字符串 ---")
	// FormatBool(b bool) string
	// 返回 "true" 或 "false"
	boolT := true
	strBoolT := strconv.FormatBool(boolT)
	fmt.Printf("strconv.FormatBool(%t) = \"%s\"\n", boolT, strBoolT)

	boolF := false
	strBoolF := strconv.FormatBool(boolF)
	fmt.Printf("strconv.FormatBool(%t) = \"%s\"\n", boolF, strBoolF)

	fmt.Println("\n--- strconv 包学习结束 ---")
}
