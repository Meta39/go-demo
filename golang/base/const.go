package main

import "fmt"

/*
相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。
*/

// 声明了pi和e这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。
const pi = 3.1415
const e = 2.7182

// 多个常量也可以一起声明：
const (
	pi2 = 3.1415
	e2  = 2.7182
)

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：
const (
	n1 = 100
	n2
	n3
)

/*
iota是go语言的常量计数器，只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
*/
const (
	n4 = iota //0
	n5        //1
	n6        //2
	n7        //3
)

// 几个常见的iota示例:
// 1.使用_跳过某些值
const (
	jump  = iota //0
	jump2        //1
	_
	jump4 //3
)

// 2.iota声明中间插队
const (
	cut1 = iota //0
	cut2 = 100  //100
	cut3 = iota //2
	cut4        //3
)
const cut5 = iota //0

// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 多个iota定义在一行
const (
	a, b  = iota + 1, iota + 2 //1,2
	c, d                       //2,3
	e1, f                      //3,4
)

func Const() {
	fmt.Println("============ 常量 ============")
	fmt.Println(pi, e)
	fmt.Println(pi2, e2)
	fmt.Println(n1, n2, n3)
	fmt.Println(n4, n5, n6, n7)
	fmt.Println(jump, jump2, jump4)
	fmt.Println(cut1, cut2, cut3, cut4)
	fmt.Println(cut5)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e1, f)
	fmt.Println("============ 常量 ============")
}
