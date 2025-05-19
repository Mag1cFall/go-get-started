package main

import "fmt"

// --- 1. 定义结构体 ---
// 使用 type 关键字和 struct 关键字来定义结构体。
// 结构体是一种聚合数据类型，可以将多个不同类型的字段（成员变量）组合在一起。

// 定义一个 Person 结构体
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	IsActive  bool
}

// 定义一个 Address 结构体，演示结构体嵌套
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// 定义一个 Employee 结构体，它将嵌套 Person 和 Address
// 也可以使用匿名字段（嵌入类型）
type Employee struct {
	Person      // 匿名字段 (嵌入 Person 结构体)
	Department  string
	Salary      float64
	ContactInfo Address // 命名字段，类型是 Address 结构体
}

func main() {
	fmt.Println("--- 第2周学习：结构体 (Structs) ---")

	// --- 2. 创建结构体实例 ---
	fmt.Println("\n--- 2. 创建结构体实例 ---")

	// a) 使用 var 声明，然后逐个字段赋值 (此时字段为零值)
	var p1 Person
	fmt.Println("p1 (零值初始化):", p1) // 输出: {  0  false} (字符串为空, int为0, bool为false)
	p1.FirstName = "Alice"
	p1.LastName = "Smith"
	p1.Age = 30
	p1.Email = "alice.smith@example.com"
	p1.IsActive = true
	fmt.Println("p1 (赋值后):", p1)

	// b) 使用字面量创建并初始化 (推荐)
	p2 := Person{
		FirstName: "Bob",
		LastName:  "Johnson",
		Age:       25,
		Email:     "bob.j@example.com",
		IsActive:  false,
	}
	fmt.Println("p2 (字面量初始化):", p2)

	// c) 如果按字段顺序提供所有值，可以省略字段名 (不推荐，容易出错)
	p3 := Person{"Charlie", "Brown", 35, "charlie@example.com", true}
	fmt.Println("p3 (按顺序初始化，不推荐):", p3)

	// d) 使用 new() 函数创建结构体指针
	// new(T) 返回一个 *T 指针，指向一个 T 类型的零值实例
	p4Ptr := new(Person)
	fmt.Println("p4Ptr (new创建的指针):", p4Ptr) // 输出: &{ 0  false}
	fmt.Println("*p4Ptr (解引用):", *p4Ptr)    // 输出: {  0  false}
	p4Ptr.FirstName = "Diana"               // Go 允许直接通过结构体指针访问字段 (自动解引用)
	p4Ptr.Age = 28
	fmt.Println("p4Ptr (赋值后):", p4Ptr)
	fmt.Println("*p4Ptr (赋值后):", *p4Ptr)

	// e) 获取结构体字面量的地址 (返回指针)
	p5Ptr := &Person{
		FirstName: "Eve",
		LastName:  "Adams",
		Age:       22,
	}
	fmt.Println("p5Ptr (字面量地址):", p5Ptr)

	// --- 3. 访问结构体字段 ---
	// 使用点 . 操作符访问结构体的字段
	fmt.Println("\n--- 3. 访问结构体字段 ---")
	fmt.Println("p1 的名字:", p1.FirstName, p1.LastName)
	fmt.Println("p2 的年龄:", p2.Age)
	fmt.Println("p4Ptr (指针) 的名字:", p4Ptr.FirstName) // Go 自动解引用: (*p4Ptr).FirstName

	// --- 4. 结构体作为函数参数和返回值 ---
	fmt.Println("\n--- 4. 结构体作为函数参数和返回值 ---")
	printPersonInfo(p1)

	p2Modified := updatePersonAge(p2, 26) // 结构体是值类型，传递的是副本
	fmt.Println("p2 (原始):", p2)
	fmt.Println("p2Modified (年龄更新后):", p2Modified)

	updatePersonAgeByPtr(&p3, 36) // 通过指针修改原始结构体
	fmt.Println("p3 (通过指针更新年龄后):", p3)

	// --- 5. 结构体嵌套与匿名字段 (嵌入) ---
	fmt.Println("\n--- 5. 结构体嵌套与匿名字段 ---")

	emp1 := Employee{
		Person: Person{ // 初始化嵌入的 Person
			FirstName: "Gary",
			LastName:  "Oldman",
			Age:       45,
			Email:     "gary@example.com",
			IsActive:  true,
		},
		Department: "Engineering",
		Salary:     90000.00,
		ContactInfo: Address{ // 初始化命名的 Address 字段
			Street:  "123 Tech Road",
			City:    "GoCity",
			ZipCode: "12345",
		},
	}
	fmt.Println("员工 emp1:", emp1)

	// 访问匿名字段的成员 (可以直接访问，就像是 Employee 自己的字段)
	fmt.Println("emp1 名字:", emp1.FirstName, emp1.LastName) // 直接访问 Person 的字段
	fmt.Println("emp1 年龄:", emp1.Age)
	// 也可以通过类型名访问 (如果发生命名冲突时需要)
	fmt.Println("emp1.Person.Email:", emp1.Person.Email)

	// 访问命名字段的成员
	fmt.Println("emp1 地址:", emp1.ContactInfo.Street, emp1.ContactInfo.City)

	emp1.Age = 46 // 修改嵌入结构体的字段
	emp1.ContactInfo.City = "GoLand"
	fmt.Println("修改后 emp1 年龄:", emp1.Age)
	fmt.Println("修改后 emp1 城市:", emp1.ContactInfo.City)

	// --- 6. 结构体比较 ---
	// 如果结构体的所有字段都是可比较的，那么这个结构体本身也是可比较的。
	// 可以使用 == 或 != 进行比较。
	fmt.Println("\n--- 6. 结构体比较 ---")
	personA := Person{FirstName: "A", LastName: "B", Age: 10}
	personB := Person{FirstName: "A", LastName: "B", Age: 10}
	personC := Person{FirstName: "X", LastName: "Y", Age: 20}

	fmt.Println("personA == personB:", personA == personB) // true
	fmt.Println("personA == personC:", personA == personC) // false
	fmt.Println("personA != personC:", personA != personC) // true

	// 如果结构体包含不可比较的字段（如切片、map、函数），则结构体本身不可直接比较。
	// type NonComparableStruct struct {
	// 	Name string
	// 	Tags []string // 切片是不可比较的
	// }
	// nc1 := NonComparableStruct{Name: "Test", Tags: []string{"a"}}
	// nc2 := NonComparableStruct{Name: "Test", Tags: []string{"a"}}
	// fmt.Println(nc1 == nc2) // 这会导致编译错误

	fmt.Println("\n--- 结构体学习结束 ---")
}

// printPersonInfo 接收一个 Person 结构体 (值传递)
func printPersonInfo(p Person) {
	fmt.Printf("  打印信息: %s %s, 年龄: %d, 邮箱: %s, 活跃: %t\n",
		p.FirstName, p.LastName, p.Age, p.Email, p.IsActive)
}

// updatePersonAge 接收一个 Person 结构体，返回一个新的 Person 结构体
func updatePersonAge(p Person, newAge int) Person {
	p.Age = newAge // 修改的是 p 的副本
	return p
}

// updatePersonAgeByPtr 接收一个 *Person 指针，直接修改原始结构体
func updatePersonAgeByPtr(p *Person, newAge int) {
	if p != nil {
		p.Age = newAge // 修改指针指向的原始结构体的 Age 字段
	}
}
