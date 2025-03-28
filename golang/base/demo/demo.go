package demo

import "fmt"

// 包级别标识符的可见性 同样的规则也适用于结构体，结构体中可导出字段的字段名称必须首字母大写。

// num 定义一个全局整型变量
// 首字母小写，对外不可见(只能在当前包内使用)
var num = 100

// Mode 定义一个常量
// 首字母大写，对外可见(可在其它包中使用)
const Mode = 1

func init() {
	fmt.Println("我是demo包下的全局常量num只能在demo包内使用", num)
	fmt.Println("我是demo包下的全局常量Mode，任何地方都能使用我", Mode)
	sayHi()
}

// person 定义一个代表人的结构体
// 首字母小写，对外不可见(只能在当前包内使用)
type person struct {
	name string
	Age  int
}

// Add 返回两个整数和的函数
// 首字母大写，对外可见(可在其它包中使用)
func Add(x, y int) int {
	return x + y
}

// sayHi 打招呼的函数
// 首字母小写，对外不可见(只能在当前包内使用)
func sayHi() {
	var myName = "Meta39" // 函数局部变量，只能在当前函数内使用
	fmt.Println(myName)
	p := &person{"Meta", 18}
	fmt.Println("我是demo包下的结构体person只能在demo包内使用", *p)
}
