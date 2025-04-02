package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
	//_ "github.com/lib/pq"            // PostgreSQL 驱动
	"iter"
	"log"
	"os"
	"strings"
)

/*
Iter 迭代器（核心价值：将数据生产、中间操作、最终消费分离，让代码更符合单一职责原则。）
func(yield func() bool)
func(yield func(V) bool)
func(yield func(K, V) bool)

Go 1.23 中增加的迭代器是一个函数类型，它把另一个函数（ yield 函数）作为参数，将容器中的连续元素传递给 yield 函数。
迭代器函数会在以下两种情况停止。
在序列迭代结束后停止，表示此次迭代结束。
在 yield 函数返回 false 时停止，表示提前停止迭代。
Go 标准库 iter 包 中定义了 Seq 和 Seq2 作为迭代器的简称，它们将每个序列元素的 1 或 2 个值传递给 yield 函数：

type (

	Seq[V any]     func(yield func(V) bool)
	Seq2[K, V any] func(yield func(K, V) bool)

)

其中：
一、Seq 是 sequence 的缩写，因为迭代器会循环遍历一系列值
二、Seq2 表示成对值的序列，通常是键值对或索引值对。
看到这里，你应该就明白了为什么需要改进 for/range 语句来支持单参数函数类型了，因为 for/range 语句要支持遍历新增的迭代器。
*/
func Iter() {
	fmt.Println("============ 迭代器 ============")
	/*
		Push 迭代器（标准迭代器）
		实现迭代器
		我们现在为先前定义的集合类型定义一个返回所有元素的迭代器。
	*/
	fmt.Println("标准迭代器")
	forRangeSet()

	/*
		Pull 迭代器
		我们已经了解了如何在 for/range 循环中使用迭代器。但简单的循环并不是使用迭代器的唯一方法。
		例如，有时我们可能需要并行迭代两个容器。这时我们就需要用到另外一种不同类型的迭代器：Pull 迭代器。

		Push 迭代器和 Pull 迭代器的区别：
		一、Push 迭代器将序列中的每个值推送到 yield 函数。Push 迭代器是 Go 标准库中的标准迭代器，并由 for/range 语句直接支持。
		二、Pull 迭代器的工作方式则相反。每次调用 Pull 迭代器时，它都会从序列中拉出另一个值并返回该值；
		for/range 语句不直接支持 Pull 迭代器；但可以通过编写一个简单的 for 循环遍历 Pull 迭代器。

		通常不需要自行实现一个 Pull 迭代器，新的标准库中 iter.Pull 和 iter.Pull2 函数能够将标准迭代器转为 Pull 迭代器。

			func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func())
			func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func())

		它们会返回两个函数。
		一、第一个是 Pull 迭代器：每次调用时都会返回序列中的下一个值和一个布尔值，该布尔值表示该值是否有效。
		二、第二个是停止函数，应在完成 Pull 迭代器后调用。
		Pull 迭代器示例，将一个迭代器中的两个连续值对作为一个元素，返回一个新的迭代器。
	*/
	fmt.Println("Pull 迭代器")
	fmt.Println("Pull 迭代器基础用法")
	// 示例1: 基础用法
	seq := SliceToSeq([]int{1, 2, 3, 4})
	for a, b := range Pairs(seq) {
		fmt.Printf("(%d, %d)\n", a, b)
	}

	fmt.Println("Pull 迭代器奇数长度序列")
	// 示例2: 奇数长度序列
	seq2 := SliceToSeq([]string{"a", "b", "c"})
	for a, b := range Pairs(seq2) {
		fmt.Printf("(%q, %q)\n", a, b)
	}

	fmt.Println("Pull 迭代器空序列")
	// 示例3: 空序列
	seq3 := SliceToSeq([]float64{})
	for a, b := range Pairs(seq3) {
		fmt.Printf("(%f, %f)\n", a, b) // 不会执行
	}

	fmt.Println("iter替代for的应用场景")
	// 如果用for调用方必须处理错误
	fmt.Println("场景1：for逐行读取文件并处理（缺点：调用方必须处理错误）")
	err := processFileWithForLoop("log.txt")
	if err != nil {
		//处理错误
		fmt.Println(err)
		//return
	}

	fmt.Println("场景1：iter逐行读取文件并处理（优点：【生产者：只管读取文件，不关心数据如何被使用】）")
	fmt.Println("消费者1：过滤错误日志")
	// 消费者1：过滤错误日志
	for line := range processFileWithIterLoop("log.txt") {
		if strings.HasPrefix(line, "ERROR") {
			fmt.Println(line)
		}
	}

	fmt.Println("消费者2：统计行数（复用processFileWithIterLoop）")
	// 消费者2：统计行数（复用processFileWithIterLoop）
	count := 0
	for range processFileWithIterLoop("log.txt") {
		count++
	}

	fmt.Println("场景2：生成斐波那契数列")
	fmt.Println("传统 for 循环实现生成斐波那契数列")
	printFibonacciWithForLoop(1)
	fmt.Println("iter实现生成斐波那契数列")
	fmt.Println("用法1：打印前10项")
	// 用法1：打印前10项
	for v := range printFibonacciWithIterLoop() {
		// 消费者控制退出
		if v > 100 {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("用法2：求和（同一迭代器不同消费方式）")
	// 用法2：求和（同一迭代器不同消费方式）
	sum := 0
	for v := range printFibonacciWithIterLoop() {
		// 消费者控制退出
		if v > 1000 {
			break
		}
		sum += v
	}
	fmt.Printf("sum:%v\n", sum)

	fmt.Println("场景3：数据库查询")
	// 1. 创建并连接数据库
	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 确保关闭连接
	fmt.Println("传统 for 循环实现数据库查询")
	fmt.Println("iter实现数据库查询")
	queryUsersWithForLoop(db)
	fmt.Println("消费者1：打印数据")
	// 消费者1：打印数据
	for id, name := range queryUsersWithIterLoop(db) {
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	fmt.Println("消费者2：转换为Map（queryUsersWithIterLoop）")
	// 消费者2：转换为Map（queryUsersWithIterLoop）
	users := make(map[int]string)
	for id, name := range queryUsersWithIterLoop(db) {
		users[id] = name
	}

	fmt.Println("============ 迭代器 ============")
}

// Set 基于 map 定义一个存放元素的集合类型
type Set[E comparable] struct {
	m map[E]struct{}
}

// NewSet 返回一个 set
func NewSet[E comparable]() *Set[E] {
	return &Set[E]{m: make(map[E]struct{})}
}

// Add 向 set 中添加元素
func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

// All 返回一个迭代器，迭代集合中的所有元素
func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

// 使用迭代器All。
func forRangeSet() {
	s := NewSet[string]()
	s.Add("Golang")
	s.Add("Java")
	s.Add("Python")
	s.Add("C++")

	//当我们调用 s.All 后会得到一个迭代器函数，然后可以直接使用 for/range 语句来遍历它。
	for v := range s.All() {
		fmt.Println(v)
	}
}

// SliceToSeq 辅助函数：将切片转换为Seq
func SliceToSeq[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

// Pairs 返回一个迭代器，遍历 seq 中连续的值对。
func Pairs[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v1, ok1 := next()
			if !ok1 {
				return
			}
			v2, ok2 := next()
			// If ok2 is false, v2 should be the
			// zero value; yield one last pair.
			if !yield(v1, v2) {
				return
			}
			if !ok2 {
				return
			}
		}
	}
}

//================================================= iter替代for应用场景 =================================================

/*
关键结论
场景			for 循环痛点					迭代器解决方案
数据来源变化	需重写整个循环（如文件→网络）		只需替换生产者（ReadLines→ReadNet）
组合操作		嵌套循环导致代码混乱			通过 Filter/Map 等组合迭代器
资源管理		需手动调用 Close()			通过 defer 自动管理
无限数据流	无法实现						自然支持（如 printFibonacciWithIterLoop()）

迭代器的核心价值：将数据生产、中间操作、最终消费分离，让代码更符合单一职责原则。
*/

/*
processFileWithForLoop
场景1：逐行读取文件并处理
问题：
一、消费逻辑（过滤和打印）与文件操作耦合
二、无法复用文件读取逻辑（如想改成网络流需重写）
*/
func processFileWithForLoop(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ERROR") { // 过滤条件
			fmt.Println(line)
		}
	}
	return scanner.Err()
}

/*
processFileWithIterLoop
场景1：iter逐行读取文件并处理
优势：
1、文件读取逻辑可复用
2、消费逻辑可以自由组合（如再加一个 Filter 迭代器）
*/
func processFileWithIterLoop(file string) iter.Seq[string] {
	return func(yield func(string) bool) {
		f, err := os.Open(file)
		if err != nil {
			log.Println("open failed:", err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if !yield(scanner.Text()) { // 将控制权交给消费者
				break
			}
		}
	}
}

/*
printFibonacciWithForLoop 场景2：生成斐波那契数列
传统 for 循环实现
问题：
1、硬编码了输出逻辑（fmt.Println）
2、无法实现"无限序列"（必须提前知道 n）
*/
func printFibonacciWithForLoop(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

/*
printFibonacciWithIterLoop 场景2：生成斐波那契数列
迭代器实现
优势：
1、支持无限序列
2、生产逻辑与消费逻辑解耦
*/
func printFibonacciWithIterLoop() iter.Seq[int] {
	a, b := 0, 1
	return func(yield func(int) bool) {
		for {
			if !yield(a) { // 允许消费者控制终止
				return
			}
			a, b = b, a+b
		}
	}
}

/*
创建数据库并插入测试数据
sql.Open() 时通过连接字符串前缀（如 "mysql://..." 或 "postgres://..."）动态选择驱动，只有实际被使用的驱动会被激活。
因此引入的驱动都可以用 _ 进行匿名。这种设计正是 Go “显式依赖 + 按需加载” 哲学的优秀体现。
*/
func createDatabase() (*sql.DB, error) {
	// 1.连接 MySQL（只会用到 go-sql-driver/mysql）
	dsn := "root:123456@tcp(localhost:3306)/db_go?charset=utf8mb4&parseTime=true&tls=skip-verify&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %v", err)
	}

	// 连接 PostgreSQL（只会用到 lib/pq）
	//pgDB, _ := sql.Open("postgres", "postgres://user:pass@localhost/pgdb?sslmode=disable")

	// 2.创建表
	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name VARCHAR(100)
        )`); err != nil {
		return nil, fmt.Errorf("建表失败: %v", err)
	}

	// 3.清空旧数据（可选）【生成环境不要这么干！！！】
	if _, err := db.Exec("DELETE FROM users"); err != nil {
		return nil, fmt.Errorf("清空数据失败: %v", err)
	}

	// 4.插入测试数据
	if _, err := db.Exec(`
        INSERT INTO users (id, name) VALUES
            (1, 'Alice'),
            (2, 'Bob'),
            (3, 'Charlie')`); err != nil {
		return nil, fmt.Errorf("插入数据失败: %v", err)
	}

	return db, nil
}

/*
queryUsersWithForLoop 场景3：数据库查询
传统 for 循环实现
问题：
1、业务逻辑（如打印）与资源管理耦合
2、无法复用查询逻辑（如想改成JSON输出需重写）
*/
func queryUsersWithForLoop(db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		//处理错误
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			//处理错误
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}

/*
queryUsersWithIterLoop 场景3：数据库查询
迭代器实现
优势：
1、查询逻辑只需写一次
2、资源（如 rows.Close()）自动管理
*/
func queryUsersWithIterLoop(db *sql.DB) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				return
			}
			if !yield(id, name) {
				break
			} // 将数据交给消费者
		}
	}
}
