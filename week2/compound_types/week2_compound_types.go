package main

import "fmt"

func main() {
	fmt.Println("--- 第2周学习：复合类型 (数组、切片、Map) ---")

	// --- 1. 数组 (array) ---
	// 数组是固定长度的、同一类型元素的集合。长度是其类型的一部分。
	fmt.Println("\n--- 1. 数组 (array) ---")

	// 声明一个包含5个整数的数组，默认值为0
	var numbers [5]int
	fmt.Println("默认初始化的数组:", numbers) // 输出: [0 0 0 0 0]

	// 给数组元素赋值
	numbers[0] = 10
	numbers[1] = 20
	numbers[4] = 50
	fmt.Println("赋值后的数组:", numbers) // 输出: [10 20 0 0 50]
	fmt.Println("数组的第一个元素:", numbers[0])
	fmt.Println("数组的长度:", len(numbers)) // len() 用于获取数组、切片、字符串、map或channel的长度

	// 声明并初始化数组
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("初始化的字符串数组:", names)

	// 使用 ... 自动推断数组长度
	languages := [...]string{"Go", "Python", "JavaScript"}
	fmt.Println("自动推断长度的数组:", languages, "长度:", len(languages))

	// 多维数组
	var matrix [2][3]int // 2行3列的整数数组
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][2] = 6
	fmt.Println("多维数组:", matrix)

	// 遍历数组
	fmt.Println("遍历 names 数组:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("索引 %d: %s\n", i, names[i])
	}

	fmt.Println("使用 range 遍历 languages 数组:")
	// for...range 循环可以用于遍历数组、切片、字符串、map 和 channel
	// 对于数组和切片，range 返回索引和对应元素的值
	for index, value := range languages {
		fmt.Printf("索引 %d: 值 %s\n", index, value)
	}
	// 如果你不需要索引，可以用下划线 _ 忽略它
	for _, value := range languages {
		fmt.Printf("值: %s\n", value)
	}

	// --- 2. 切片 (slice) ---
	// 切片是对底层数组一个连续片段的引用（或称视图）。切片是动态大小的。
	// 切片比数组更常用，因为它们更灵活。
	fmt.Println("\n--- 2. 切片 (slice) ---")

	// 声明切片 (未初始化时为 nil)
	var scores []int
	fmt.Println("声明的空切片:", scores, "是否为nil?", scores == nil) // 输出: [] true

	// 使用 make 创建切片
	// make([]T, length, capacity)
	// length: 切片的初始长度
	// capacity: 切片底层数组的容量 (可选，默认等于length)
	ages := make([]int, 3, 5)                                             // 长度为3，容量为5的int切片
	fmt.Println("使用make创建的切片:", ages, "长度:", len(ages), "容量:", cap(ages)) // 输出: [0 0 0] 3 5
	ages[0] = 30
	ages[1] = 25
	// ages[3] = 40 // 这会引发 panic: runtime error: index out of range [3] with length 3 (因为长度是3)

	// 使用字面量初始化切片 (类似数组，但不指定长度)
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("字面量初始化的切片:", fruits, "长度:", len(fruits), "容量:", cap(fruits))

	// 从数组创建切片 (切片表达式 a[low:high])
	// 这是一个半开区间，包括 low，不包括 high
	primesArray := [...]int{2, 3, 5, 7, 11, 13, 17}
	subPrimes := primesArray[1:4]                                       // 从索引1到索引3 (不包括4)
	fmt.Println("从数组创建的切片:", subPrimes)                                 // 输出: [3 5 7]
	fmt.Println("subPrimes 长度:", len(subPrimes), "容量:", cap(subPrimes)) // 容量会从切片开始位置到底层数组末尾

	allPrimes := primesArray[:]       // 包含所有元素
	firstThree := primesArray[:3]     // 从开头到索引2
	afterIndexFour := primesArray[4:] // 从索引4到末尾
	fmt.Println("allPrimes:", allPrimes, "firstThree:", firstThree, "afterIndexFour:", afterIndexFour)

	// 切片是引用类型：修改切片会影响底层数组和其他引用同一数组的切片
	fmt.Println("原始数组 primesArray:", primesArray)
	subPrimes[0] = 333                                      // 修改切片元素
	fmt.Println("修改 subPrimes 后，primesArray:", primesArray) // 底层数组被修改
	fmt.Println("修改 subPrimes 后，subPrimes:", subPrimes)

	// 使用 append 向切片添加元素
	// 如果切片的容量不足，append 会创建一个新的更大的底层数组，并将元素复制过去
	var emptySlice []int
	emptySlice = append(emptySlice, 1)
	fmt.Println("append 到空切片:", emptySlice, "len:", len(emptySlice), "cap:", cap(emptySlice))
	emptySlice = append(emptySlice, 2, 3, 4)
	fmt.Println("append 多个元素:", emptySlice, "len:", len(emptySlice), "cap:", cap(emptySlice))

	slice1 := []string{"a", "b"}
	slice2 := []string{"c", "d", "e"}
	slice1 = append(slice1, slice2...) // 使用 ... 将一个切片的所有元素追加到另一个切片
	fmt.Println("append 另一个切片:", slice1)

	// copy 函数：用于复制切片内容
	// copy(dst, src) 返回复制的元素数量
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	numCopied := copy(dst, src)
	fmt.Println("源切片 src:", src)
	fmt.Println("目标切片 dst:", dst, "(复制了", numCopied, "个元素)")
	dst[0] = 100 // 修改 dst 不会影响 src，因为它们引用不同的底层数组 (make创建了新的)
	fmt.Println("修改 dst 后, src:", src, "dst:", dst)

	// 遍历切片 (与数组类似)
	fmt.Println("遍历 fruits 切片:")
	for i, fruit := range fruits {
		fmt.Printf("索引 %d: %s\n", i, fruit)
	}

	// --- 3. Map (映射) ---
	// Map 是一种无序的键值对集合。键必须是可比较的类型，值可以是任意类型。
	fmt.Println("\n--- 3. Map ---")

	// 声明 Map (未初始化时为 nil)
	var person map[string]string
	fmt.Println("声明的空map:", person, "是否为nil?", person == nil) // 输出: map[] true
	// person["name"] = "Roo" // 对nil map写入会导致panic

	// 使用 make 创建 Map
	agesMap := make(map[string]int)
	agesMap["Alice"] = 30
	agesMap["Bob"] = 25
	fmt.Println("使用make创建的map:", agesMap)

	// 使用字面量初始化 Map
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"China":  "Beijing", // 最后的逗号是允许的，甚至是推荐的（方便多行添加）
	}
	fmt.Println("字面量初始化的map:", capitals)
	fmt.Println("日本的首都是:", capitals["Japan"])

	// 访问 Map 元素
	// 如果键不存在，会返回对应值类型的零值
	unknownCapital := capitals["Germany"]
	fmt.Println("德国的首都是 (不存在):", unknownCapital) // 输出空字符串

	// 判断键是否存在
	// value, ok := myMap[key]
	capitalOfGermany, ok := capitals["Germany"]
	if ok {
		fmt.Println("德国的首都是:", capitalOfGermany)
	} else {
		fmt.Println("德国的首都信息未找到。")
	}

	capitalOfFrance, ok := capitals["France"]
	if ok {
		fmt.Println("法国的首都是:", capitalOfFrance)
	} else {
		fmt.Println("法国的首都信息未找到。")
	}

	// 添加或修改 Map 元素
	capitals["Germany"] = "Berlin" // 添加新键值对
	capitals["France"] = "PARIS"   // 修改已存在的键的值
	fmt.Println("修改和添加后的map:", capitals)

	// 删除 Map 元素
	// delete(myMap, key)
	delete(capitals, "Japan")
	fmt.Println("删除日本后的map:", capitals)
	delete(capitals, "USA") // 删除不存在的键不会报错

	// 获取 Map 长度 (键值对的数量)
	fmt.Println("capitals map 的长度:", len(capitals))

	// 遍历 Map
	// 注意：Map 的遍历顺序是不确定的
	fmt.Println("遍历 agesMap:")
	for name, age := range agesMap {
		fmt.Printf("%s 的年龄是 %d\n", name, age)
	}

	// 只遍历键
	fmt.Println("只遍历 capitals 的键:")
	for country := range capitals {
		fmt.Println("国家:", country)
	}

	fmt.Println("\n--- 复合类型学习结束 ---")
}
