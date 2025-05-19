package main

import (
	"fmt"
	"strings" // 导入 strings 包
)

func main() {
	fmt.Println("--- 第4周学习：常用标准库 (strings 包) ---")

	s := "Hello, Go World! Go is Awesome. Go Go Go!"
	fmt.Printf("原始字符串: \"%s\"\n", s)

	// --- 1. 检查包含关系 ---
	fmt.Println("\n--- 1. 检查包含关系 ---")
	// Contains(s, substr string) bool: 判断字符串 s 是否包含子串 substr
	fmt.Printf("strings.Contains(\"%s\", \"World\"): %t\n", s, strings.Contains(s, "World"))
	fmt.Printf("strings.Contains(\"%s\", \"world\"): %t\n", s, strings.Contains(s, "world")) // 区分大小写

	// ContainsAny(s, chars string) bool: 判断字符串 s 是否包含 chars 中的任意一个字符
	fmt.Printf("strings.ContainsAny(\"%s\", \"Wxyz\"): %t\n", s, strings.ContainsAny(s, "Wxyz")) // 包含 W
	fmt.Printf("strings.ContainsAny(\"%s\", \"xyz\"): %t\n", s, strings.ContainsAny(s, "xyz"))   // 不包含

	// ContainsRune(s string, r rune) bool: 判断字符串 s 是否包含符文 r
	fmt.Printf("strings.ContainsRune(\"%s\", 'o'): %t\n", s, strings.ContainsRune(s, 'o')) // 'o' 是一个 rune

	// --- 2. 计数 ---
	fmt.Println("\n--- 2. 计数 ---")
	// Count(s, substr string) int: 计算子串 substr 在字符串 s 中出现的次数 (非重叠)
	fmt.Printf("strings.Count(\"%s\", \"Go\"): %d\n", s, strings.Count(s, "Go"))
	fmt.Printf("strings.Count(\"%s\", \"o\"): %d\n", s, strings.Count(s, "o"))

	// --- 3. 前缀和后缀 ---
	fmt.Println("\n--- 3. 前缀和后缀 ---")
	// HasPrefix(s, prefix string) bool: 判断字符串 s 是否以 prefix 开头
	fmt.Printf("strings.HasPrefix(\"%s\", \"Hello\"): %t\n", s, strings.HasPrefix(s, "Hello"))
	// HasSuffix(s, suffix string) bool: 判断字符串 s 是否以 suffix 结尾
	fmt.Printf("strings.HasSuffix(\"%s\", \"Go!\"): %t\n", s, strings.HasSuffix(s, "Go!"))

	// --- 4. 查找索引 ---
	fmt.Println("\n--- 4. 查找索引 ---")
	// Index(s, substr string) int: 返回子串 substr 在字符串 s 中第一次出现的索引，如果不存在则返回 -1
	fmt.Printf("strings.Index(\"%s\", \"Go\"): %d\n", s, strings.Index(s, "Go"))         // 第一个 Go
	fmt.Printf("strings.Index(\"%s\", \"Python\"): %d\n", s, strings.Index(s, "Python")) // 不存在

	// LastIndex(s, substr string) int: 返回子串 substr 在字符串 s 中最后一次出现的索引
	fmt.Printf("strings.LastIndex(\"%s\", \"Go\"): %d\n", s, strings.LastIndex(s, "Go")) // 最后一个 Go

	// IndexAny(s, chars string) int: 返回 chars 中任意字符在 s 中首次出现的索引
	fmt.Printf("strings.IndexAny(\"%s\", \"xyzW\"): %d\n", s, strings.IndexAny(s, "xyzW")) // W 的索引

	// --- 5. 分割字符串 ---
	fmt.Println("\n--- 5. 分割字符串 ---")
	// Split(s, sep string) []string: 将字符串 s 按照分隔符 sep 分割成一个字符串切片
	sentence := "The quick brown fox"
	words := strings.Split(sentence, " ")
	fmt.Printf("strings.Split(\"%s\", \" \"): %v (类型: %T)\n", sentence, words, words)
	for i, word := range words {
		fmt.Printf("  词 %d: %s\n", i, word)
	}

	// SplitN(s, sep string, n int) []string: 最多分割 n 次，返回的切片最多有 n 个元素
	// 如果 n < 0，则等同于 Split
	data := "apple,banana,cherry,date"
	partsN := strings.SplitN(data, ",", 3) // 最多分出3个部分
	fmt.Printf("strings.SplitN(\"%s\", \",\", 3): %v\n", data, partsN)

	// --- 6. 连接字符串 ---
	fmt.Println("\n--- 6. 连接字符串 ---")
	// Join(a []string, sep string) string: 将字符串切片 a 中的所有元素用分隔符 sep 连接成一个字符串
	elements := []string{"Go", "is", "fun"}
	joinedString := strings.Join(elements, "-")
	fmt.Printf("strings.Join(%v, \"-\"): \"%s\"\n", elements, joinedString)

	// --- 7. 大小写转换 ---
	fmt.Println("\n--- 7. 大小写转换 ---")
	mixedCase := "Go Is FuN"
	fmt.Printf("原始: \"%s\"\n", mixedCase)
	// ToLower(s string) string: 转换为小写
	fmt.Printf("strings.ToLower: \"%s\"\n", strings.ToLower(mixedCase))
	// ToUpper(s string) string: 转换为大写
	fmt.Printf("strings.ToUpper: \"%s\"\n", strings.ToUpper(mixedCase))
	// ToTitle(s string) string: 将每个单词的首字母大写 (更推荐使用 cases.Title，见下)
	fmt.Printf("strings.ToTitle (旧，推荐用golang.org/x/text/cases): \"%s\"\n", strings.ToTitle(mixedCase))
	// 注意：strings.ToTitle 已经被标记为不推荐，因为它不能很好地处理 Unicode。
	// 推荐使用 golang.org/x/text/cases 和 golang.org/x/text/language 来进行更准确的大小写转换。
	// (这需要引入外部模块，我们暂时只关注标准库 strings)

	// --- 8. 替换 ---
	fmt.Println("\n--- 8. 替换 ---")
	// Replace(s, old, new string, n int) string:
	// 将字符串 s 中的前 n 个不重叠的 old 子串替换为 new 子串。
	// 如果 n < 0，则替换所有匹配的子串。
	fmt.Printf("strings.Replace(\"%s\", \"Go\", \"Golang\", 1): \"%s\"\n", s, strings.Replace(s, "Go", "Golang", 1))
	fmt.Printf("strings.Replace(\"%s\", \"Go\", \"Golang\", 2): \"%s\"\n", s, strings.Replace(s, "Go", "Golang", 2))
	fmt.Printf("strings.ReplaceAll(\"%s\", \"Go\", \"Golang\"): \"%s\"\n", s, strings.ReplaceAll(s, "Go", "Golang")) // 等价于 n = -1

	// --- 9. 去除空白 ---
	fmt.Println("\n--- 9. 去除空白 ---")
	spacedString := "  \t Hello, Spaces! \n  "
	fmt.Printf("带空白的字符串: \"%s\"\n", spacedString)
	// TrimSpace(s string) string: 去除字符串 s 两端的空白字符 (空格, \t, \n, \r 等)
	fmt.Printf("strings.TrimSpace: \"%s\"\n", strings.TrimSpace(spacedString))
	// Trim(s string, cutset string) string: 去除字符串 s 两端包含在 cutset 中的任意字符
	fmt.Printf("strings.Trim(\"¡¡¡Hello!!!\", \"¡!\"): \"%s\"\n", strings.Trim("¡¡¡Hello!!!", "¡!"))
	// TrimLeft(s string, cutset string) string: 去除左端
	// TrimRight(s string, cutset string) string: 去除右端
	fmt.Printf("strings.TrimLeft(\"___Hello\", \"_\"): \"%s\"\n", strings.TrimLeft("___Hello", "_"))
	fmt.Printf("strings.TrimRight(\"Hello___\", \"_\"): \"%s\"\n", strings.TrimRight("Hello___", "_"))

	// --- 10. 字符串构建器 strings.Builder ---
	// 对于需要多次拼接字符串的场景，直接使用 + 或 += 会导致多次内存分配和复制，效率较低。
	// strings.Builder 类型提供了一种更高效的方式来构建字符串。
	fmt.Println("\n--- 10. strings.Builder ---")
	var builder strings.Builder
	builder.WriteString("这是一个")
	builder.WriteByte(' ') // 写入单个字节
	builder.WriteString("字符串构建器")
	builder.WriteRune('。')          // 写入单个 rune (Unicode 字符)
	finalString := builder.String() // 获取最终构建的字符串
	fmt.Println("使用 strings.Builder 构建的字符串:", finalString)
	fmt.Println("Builder 的当前长度:", builder.Len())
	// builder.Reset() // 可以重置 Builder 以便复用

	fmt.Println("\n--- strings 包学习结束 ---")
}
