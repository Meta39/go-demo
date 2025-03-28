package main

import "fmt"

// 全局变量m
var m = 100

// Variable 2、变量
func Variable() {
	fmt.Println("============ 变量 ============")
	/*
		标准声明：var 变量名 变量类型
		变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号。
	*/
	var name string
	var age int
	var isOk bool

	fmt.Println(name, age, isOk)

	//批量声明
	var (
		a string
		b int
		c bool
		d float32
	)

	fmt.Println(a, b, c, d)

	/*
		变量的初始化：var 变量名 类型 = 表达式
		Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。
		每个变量会被初始化成其类型的默认值，例如：
		整型和浮点型变量的默认值为0。
		字符串变量的默认值为空字符串。
		布尔型变量默认为false。
		切片、函数、指针变量的默认为nil。
	*/
	var name2 string = "Meta39"
	var age2 int = 18
	fmt.Println(name2, age2)

	//或者一次初始化多个变量
	var name3, age3 = "Meta39", 20
	fmt.Println(name3, age3)

	//类型推导。有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。
	var name4 = "Meta39"
	var age4 = 18
	fmt.Println(name4, age4)

	//短变量声明。在函数内部，可以使用更简略的 := 方式声明并初始化变量。
	n := 10
	fmt.Println("使用全局变量m=100", m) //使用全局变量m
	m := 200                      //此处声明局部变量m
	fmt.Println("使用全局变量m=200", m, n)

	/*
		匿名变量。在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示，如下所示：
		匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在Lua等编程语言里，匿名变量也被叫做哑元变量。)
		1、函数外的每个语句都必须以关键字开始（var、const、func等）
		2、:=不能使用在函数外。
		3、_多用于占位，表示忽略值。
	*/
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Println("============ 变量 ============")
}

// foo 函数名小写，说明只能在当前文件使用。
func foo() (int, string) {
	return 10, "Meta39"
}
