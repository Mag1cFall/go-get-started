package main

import (
	"database/sql" // Go 标准的数据库接口包
	"fmt"
	"log"
	"time"

	// 导入 MySQL 驱动程序。
	// _ (下划线) 表示我们只导入该包以使其注册自己到 database/sql，
	// 而不直接在代码中使用该包的导出名。database/sql 会通过驱动名 "mysql" 来使用它。
	_ "github.com/go-sql-driver/mysql"
)

// DSN (Data Source Name) for MySQL connection.
// 格式: username:password@protocol(address)/dbname?param=value
// 你需要根据你的 MySQL 设置修改这个字符串。
// 例如: "root:your_password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
// - testdb: 数据库名，你需要先创建它。
// - charset=utf8mb4: 推荐的字符集。
// - parseTime=True: 允许驱动将数据库中的 DATETIME/TIMESTAMP 类型解析为 time.Time。
// - loc=Local: 使用本地时区。
const dsn = "your_username:your_password@tcp(127.0.0.1:3306)/your_dbname?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
}

func main() {
	fmt.Println("--- 第7周学习：数据库操作 (MySQL 与 database/sql) ---")
	fmt.Println("!!! 重要: 请确保你已配置好 MySQL 服务器，并更新了代码中的 dsn 连接字符串。")
	fmt.Println("!!! 并且，数据库 'your_dbname' (或你指定的库名) 需要预先创建好。")

	if dsn == "your_username:your_password@tcp(127.0.0.1:3306)/your_dbname?charset=utf8mb4&parseTime=True&loc=Local" {
		log.Println("提醒: 请更新 week7_mysql_example.go 文件中的 dsn 常量以匹配你的 MySQL 配置。")
		log.Println("本示例将不会实际连接数据库，除非 dsn 被修改。")
		// return // 可以取消注释这行，在未配置DSN时直接退出
	}

	// --- 1. 打开数据库连接 ---
	// sql.Open 不会立即建立连接或验证连接参数，它只是准备一个 *sql.DB 对象。
	// 实际的连接会在第一次需要时（例如执行查询）惰性建立。
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("  [DB] sql.Open 错误: %v\n", err) // log.Fatal 会退出程序
	}
	// 使用 defer db.Close() 来确保在 main 函数结束时关闭数据库连接池。
	defer db.Close()

	// --- 2. 验证数据库连接 (Ping) ---
	// db.Ping() 用于验证与数据库的连接是否仍然存在，如果需要则建立连接。
	err = db.Ping()
	if err != nil {
		log.Fatalf("  [DB] db.Ping 错误 (无法连接到数据库，请检查DSN和MySQL服务): %v\n", err)
	}
	fmt.Println("  [DB] 成功连接到 MySQL 数据库!")

	// --- 3. 创建表 (如果不存在) ---
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL) // db.Exec 用于执行不返回行的 SQL 语句 (如 INSERT, UPDATE, DELETE, CREATE TABLE)
	if err != nil {
		log.Fatalf("  [DB] 创建 users 表失败: %v\n", err)
	}
	fmt.Println("  [DB] users 表已确保存在。")

	// --- 4. 插入数据 (INSERT) ---
	username1, email1 := "alice_gopher", "alice@example.com"
	// Exec 返回一个 sql.Result 对象，包含 LastInsertId (如果适用) 和 RowsAffected。
	result, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username1, email1)
	if err != nil {
		// 可能是唯一约束冲突等
		fmt.Printf("  [DB] 插入用户 '%s' 失败: %v\n", username1, err)
	} else {
		lastID, _ := result.LastInsertId()  // 获取自增ID
		rowsAff, _ := result.RowsAffected() // 获取影响的行数
		fmt.Printf("  [DB] 成功插入用户 '%s', ID: %d, 影响行数: %d\n", username1, lastID, rowsAff)
	}

	username2, email2 := "bob_coder", "bob@example.org"
	_, err = db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username2, email2)
	if err != nil {
		fmt.Printf("  [DB] 插入用户 '%s' 失败: %v\n", username2, err)
	} else {
		fmt.Printf("  [DB] 成功插入用户 '%s'\n", username2)
	}

	// --- 5. 查询数据 (SELECT) ---
	fmt.Println("\n  --- 查询所有用户 ---")
	// db.Query 用于执行返回多行的 SELECT 查询。它返回一个 *sql.Rows 对象。
	rows, err := db.Query("SELECT id, username, email, created_at FROM users")
	if err != nil {
		log.Fatalf("  [DB] 查询所有用户失败: %v\n", err)
	}
	defer rows.Close() // 非常重要：确保在处理完结果集后关闭 rows

	var users []User
	for rows.Next() { // rows.Next() 迭代结果集中的每一行
		var u User
		// rows.Scan() 将当前行的数据列扫描到指定的变量地址中
		// 列的顺序必须与 SELECT 语句中的列顺序一致
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt); err != nil {
			log.Printf("  [DB] 扫描行数据错误: %v\n", err)
			continue
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil { // 检查迭代过程中是否发生错误
		log.Printf("  [DB] rows.Next() 迭代错误: %v\n", err)
	}

	for _, u := range users {
		fmt.Printf("    用户: ID=%d, Username=%s, Email=%s, CreatedAt=%s\n",
			u.ID, u.Username, u.Email, u.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// 查询单行数据 (db.QueryRow)
	fmt.Println("\n  --- 查询单个用户 (bob_coder) ---")
	var bob User
	// db.QueryRow 用于执行只返回一行的 SELECT 查询。它返回一个 *sql.Row 对象。
	// *sql.Row 的 Scan 方法会在没有找到行时返回 sql.ErrNoRows 错误。
	err = db.QueryRow("SELECT id, username, email, created_at FROM users WHERE username = ?", "bob_coder").Scan(&bob.ID, &bob.Username, &bob.Email, &bob.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("  [DB] 未找到用户 bob_coder")
		} else {
			log.Printf("  [DB] 查询用户 bob_coder 失败: %v\n", err)
		}
	} else {
		fmt.Printf("    找到用户: ID=%d, Username=%s, Email=%s, CreatedAt=%s\n",
			bob.ID, bob.Username, bob.Email, bob.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// --- 6. 更新数据 (UPDATE) ---
	_, err = db.Exec("UPDATE users SET email = ? WHERE username = ?", "alice.gopher@newdomain.com", "alice_gopher")
	if err != nil {
		log.Printf("  [DB] 更新用户 alice_gopher 邮箱失败: %v\n", err)
	} else {
		fmt.Println("  [DB] 用户 alice_gopher 邮箱已更新。")
	}

	// --- 7. 使用预处理语句 (Prepared Statements) ---
	// 预处理语句可以提高性能（如果多次执行相同的SQL结构）并防止SQL注入。
	fmt.Println("\n  --- 使用预处理语句插入新用户 ---")
	stmt, err := db.Prepare("INSERT INTO users (username, email) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("  [DB] db.Prepare 错误: %v\n", err)
	}
	defer stmt.Close() // 确保语句关闭

	_, err = stmt.Exec("charlie_dev", "charlie@dev.io")
	if err != nil {
		fmt.Printf("  [DB] 预处理语句插入 charlie_dev 失败: %v\n", err)
	} else {
		fmt.Println("  [DB] 使用预处理语句成功插入 charlie_dev。")
	}
	_, err = stmt.Exec("diana_designer", "diana@design.co") // 重用预处理语句
	if err != nil {
		fmt.Printf("  [DB] 预处理语句插入 diana_designer 失败: %v\n", err)
	} else {
		fmt.Println("  [DB] 使用预处理语句成功插入 diana_designer。")
	}

	// --- 8. 事务处理 (Transactions) ---
	// 事务用于将一组SQL操作作为一个原子单元执行：要么全部成功，要么全部失败回滚。
	fmt.Println("\n  --- 事务处理示例 ---")
	tx, err := db.Begin() // 开始一个事务
	if err != nil {
		log.Fatalf("  [DB] db.Begin (开始事务) 错误: %v\n", err)
	}

	// 在事务中执行操作
	_, err = tx.Exec("INSERT INTO users (username, email) VALUES (?, ?)", "eve_manager", "eve@example.com")
	if err != nil {
		fmt.Printf("  [DB] 事务中插入 eve_manager 失败: %v - 准备回滚...\n", err)
		tx.Rollback() // 如果出错，回滚事务
	} else {
		_, err = tx.Exec("INSERT INTO users (username, email) VALUES (?, ?)", "frank_tester", "frank@example.com")
		if err != nil {
			fmt.Printf("  [DB] 事务中插入 frank_tester 失败: %v - 准备回滚...\n", err)
			tx.Rollback() // 回滚
		} else {
			err = tx.Commit() // 全部成功，提交事务
			if err != nil {
				log.Fatalf("  [DB] tx.Commit 错误: %v\n", err)
			} else {
				fmt.Println("  [DB] 事务成功提交 (eve_manager 和 frank_tester 已插入)。")
			}
		}
	}
	// 尝试插入一个会产生唯一约束冲突的用户，演示事务回滚
	tx2, err := db.Begin()
	if err != nil {
		log.Fatalf("  [DB] db.Begin (事务2) 错误: %v\n", err)
	}

	fmt.Println("  [DB] 尝试在事务2中插入重复用户 alice_gopher...")
	_, err = tx2.Exec("INSERT INTO users (username, email) VALUES (?, ?)", "alice_gopher", "alice.is.back@example.com")
	if err != nil {
		fmt.Printf("  [DB] 事务2中插入重复用户 alice_gopher 失败 (预期错误): %v\n", err)
		fmt.Println("  [DB] 正在回滚事务2...")
		tx2.Rollback()
		fmt.Println("  [DB] 事务2已回滚。")
	} else {
		fmt.Println("  [DB] 事务2中插入重复用户成功了？这不应该发生。正在提交...") // 不应该到这里
		tx2.Commit()
	}

	// --- 9. 删除数据 (DELETE) ---
	// 为了清理，可以删除一些测试数据，但要注意不要误删重要数据
	// _, err = db.Exec("DELETE FROM users WHERE username LIKE 'eve_%' OR username LIKE 'frank_%' OR username LIKE 'charlie_%' OR username LIKE 'diana_%'")
	// if err != nil {
	// 	log.Printf("  [DB] 删除测试用户失败: %v\n", err)
	// } else {
	// 	fmt.Println("  [DB] 部分测试用户已删除。")
	// }

	fmt.Println("\n--- MySQL 数据库操作学习结束 ---")
}
