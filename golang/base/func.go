package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
Func 函数
1、函数是组织好的、可重复使用的、用于执行指定任务的代码块。
2、Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。

函数定义：
Go语言中定义函数使用func关键字，具体格式如下：

	func 函数名(参数)(返回值){
	    函数体
	}

其中：
一、函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
二、参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
三、返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
四、函数体：实现指定功能的代码块。
*/

// 定义全局变量num。如果局部变量和全局变量重名，优先访问局部变量。
var num int64 = 10

func Func() {
	fmt.Println("============ 函数 ============")
	//调用求和函数.
	x := 1
	y := 2
	sum := intSum(x, y)
	//函数的参数中如果相邻变量的类型相同，则可以省略类型
	intSum2(x, y)

	fmt.Println("调用 intSum 求和函数：", x, "+", y, "=", sum)

	//调用无传递参数和返回值的函数
	fmt.Println("调用无传递参数和返回值的函数hello")
	hello()

	//可变参数
	ret1 := intSum3()
	ret2 := intSum3(10)
	ret3 := intSum3(10, 20)
	ret4 := intSum3(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	//函数多返回值
	calcSum, calcSub := calc(3, 5)
	fmt.Printf("多返回值。返回第一个为求和的值%v，返回第二个为相减的值%v\n", calcSum, calcSub)

	//函数返回值命名
	calcSum2, calcSub2 := calc2(4, 6)
	fmt.Printf("返回值命名。返回第一个为求和的值%v，返回第二个为相减的值%v\n", calcSum2, calcSub2)

	//局部变量
	testLocalVar()
	//语句块定义的变量
	testLocalVar2(1, 2)
	//for循环语句中定义的变量，也是只在for语句块中生效
	testLocalVar3()

	fmt.Println("访问全局变量num", num)
	num := 100 //与全局变量重名
	fmt.Println("全局变量num和局部变量num同时存在时，优先使用局部变量。", num)

	//定义函数类型
	fmt.Println("定义函数类型")
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f

	//高阶函数
	fmt.Println("高阶函数")
	result := calc3(10, 5, func(a, b int) int {
		return a + b
	})
	fmt.Println(result) // 输出: 15

	//高阶函数-获取加法函数
	addFunc, err := do("+")
	if err != nil {
		fmt.Println("获取加法函数错误:", err)
		return
	}
	result2 := addFunc(5, 3)
	fmt.Println("5 + 3 =", result2) // 输出: 5 + 3 = 8

	// 高阶函数-获取减法函数
	subFunc, err := do("-")
	if err != nil {
		fmt.Println("获取减法函数错误:", err)
		return
	}

	result3 := subFunc(5, 3)
	fmt.Println("5 - 3 =", result3) // 输出: 5 - 3 = 2

	// 高阶函数-测试错误情况(输入一个不存在的函数操作符)
	_, err = do("*")
	if err != nil {
		fmt.Println("获取乘法函数错误:", err) // 输出: 错误: 无法识别的操作符
	}

	//高阶函数-直接使用返回的函数(推荐)
	if op, err := do("+"); err != nil {
		//错误处理
		fmt.Println("错误:", err)
	} else {
		//业务处理
		fmt.Println("5 + 3 =", op(5, 3)) // 输出: 5 + 3 = 8
	}

	//匿名函数
	noNameFunc()

	//闭包
	fmt.Println("闭包")
	fmt.Println("变量closePackage是一个函数并且它引用了其外部作用域中的x变量，此时closePackage就是一个闭包。 在closePackage的生命周期内，变量x也一直有效。")
	var closePackage = adder()
	fmt.Println(closePackage(10)) //10
	fmt.Println(closePackage(20)) //30
	fmt.Println(closePackage(30)) //60

	closePackage2 := adder() //属于新的引用环境，所以不会使用上面的closePackage变量名的值
	fmt.Println("属于新的引用环境，所以不会使用上面的变量closePackage的x值")
	fmt.Println(closePackage2(40)) //40
	fmt.Println(closePackage2(50)) //90

	//闭包-进阶用法1
	fmt.Println("闭包-进阶用法1")
	closePackage3 := adder2(1)
	fmt.Println(closePackage3(2))

	closePackage4 := adder2(3)
	fmt.Println(closePackage4(4))

	//闭包-进阶用法2
	fmt.Println("闭包-进阶用法2")
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("a"))    //a.txt

	//闭包-进阶用法3
	fmt.Println("闭包-进阶用法3")
	sumCalc, subCalc := closePackageCalc(0)
	/*
		因为 closePackageCalc 先执行 add 函数，再执行 sub 函数。
		1、所以 closePackageCalc 函数传入 0 ，先进行加法运算 0 + 1 得到 1 返回结果给 base，此时 base = 1
		2、得到 add 函数返回的 base 结果 1，再进行减法运算 1 - 1 得到 0 返回结果给 base。
		3、所以 sumCalc(1) = 1 ， subCalc(1) = 0
	*/
	fmt.Println(sumCalc(1), subCalc(1))

	//defer语句
	deferUsed()
	//defer经典案例
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	/*
		内置函数panic/recover（不推荐用于常规错误处理，对于预期可能发生的错误，使用error返回值）
		Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
		panic可以在任何地方引发，但recover只有在defer调用的函数中有效。
		注意事项：
		1、recover()必须搭配defer使用。
		2、defer一定要在可能引发panic的语句之前定义。
	*/
	funcA()
	funcB() //存在 panic 和 recover，如果只使用 panic，而不使用 recover。则 funcC() 不会执行。
	funcC()
	//如果为true，则抛出异常。如果为false，则不抛出异常。
	result4, err2 := funcD(false)
	if err2 != nil {
		fmt.Println(err2)
		return //这里需要让当前函数后面的代码不往下执行。因为后面的程序是正常逻辑的处理。
	}
	fmt.Println(result4)

	fmt.Println("============ 函数 ============")
}

// 定义一个求两个数之和的函数.调用有返回值的函数时，可以不接收其返回值。
func intSum(x int, y int) int {
	return x + y
}

// 函数的参数和返回值都是可选的，例如我们可以实现一个既不需要参数也没有返回值的函数
func hello() {
	fmt.Println("Hello Go")
}

// 函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：
func intSum2(x, y int) int {
	//上面的代码中，intSum函数有两个参数，这两个参数的类型均为int，因此可以省略x的类型，因为y后面有类型说明，x参数也是该类型。
	return x + y
}

/*
可变参数。
1、可变参数是指函数的参数数量不固定。
2、Go语言中的可变参数通过在参数名后加...来标识。注意：可变参数通常要作为函数的最后一个参数。
3、本质上，函数的可变参数是通过切片来实现的。
*/
func intSum3(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		//for v := range x {//如果用不上索引，也可以写成这样
		sum = sum + v
	}
	return sum
}

// 多返回值。Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名。函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 局部变量.局部变量又分为两种： 函数内定义的变量无法在该函数外使用，例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x
func testLocalVar() {
	fmt.Println("局部变量")
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

// 语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。
func testLocalVar2(x, y int) {
	fmt.Println("语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。")
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}

// for循环语句中定义的变量，也是只在for语句块中生效
func testLocalVar3() {
	fmt.Println("for循环语句中定义的变量，也是只在for语句块中生效")
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}

/*
定义函数类型(有点像java里的接口)
我们可以使用type关键字来定义一个函数类型，具体格式如下：
type calculation func(int, int) int
上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。
add和sub都能赋值给calculation类型的变量。
*/
type calculation func(int, int) int

// 就像是实现了calculation接口一样
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

/*
高阶函数
1、函数作为参数
2、函数作为返回值
*/
//函数作为参数
func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

/*
匿名函数和闭包
1、函数可以作为返回值，Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。
2、匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数
3、匿名函数多用于实现回调函数和闭包。
匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

	func(参数)(返回值){
	    函数体
	}
*/
func noNameFunc() {
	fmt.Println("匿名函数")
	// 将匿名函数保存到变量
	addNoNameFunc := func(x, y int) {
		fmt.Println(x + y)
	}
	addNoNameFunc(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

/*
闭包
闭包指的是一个函数和与其相关的引用环境组合而成的实体。
简单来说，闭包=函数+引用环境。
闭包其实并不复杂，只要牢记闭包=函数+引用环境。
*/
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包-进阶用法1
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包-进阶用法2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包-进阶用法3
func closePackageCalc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

/*
defer语句
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
1、在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
2、由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
3、defer执行时机：在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
|———————————————————————————————|——————————————————————————————|
|	函数中return语句底层实现		|		defer执行时机			   |
|			  	  返回值 = x		|					返回值 = x  |
|	return x   →	↓			|						↓	   |
| 				  RET指令		|	return x   →	运行defer   |
|								|						↓	   |
|								|					RET指令	   |
|———————————————————————————————|——————————————————————————————|
*/
func deferUsed() {
	fmt.Println("defer语句的用法。由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。")
	fmt.Println("==== 1、defer 开始 ====")
	defer fmt.Println("5、我最后执行")
	defer fmt.Println("4、我其次执行")
	defer fmt.Println("3、我最先执行")
	fmt.Println("==== 2、defer 结束 ====")
}

// 因为f1()的返回值不是变量，所以返回的值是1，而不是2
func f1() int {
	x := 1
	defer func() {
		x++
		fmt.Println("2、defer f1() x=", x)
	}()

	fmt.Println("1、f1() x=", x)
	return x
}

// 因为f2()的返回值是变量x，所以返回值是3
func f2() (x int) {
	defer func() {
		//拿到f2()的返回值2，在进行 + 1，得到x = 3
		fmt.Println("2、defer f2() 得到 x=", x)
		x++
		fmt.Println("3、defer f2() x=", x)
	}() //因为()没有传值，因此不存在拷贝，直接使用的是x = 2的值
	fmt.Println("1、f2() x=", x) //x =0
	return 2
	//x = 2
}

func f3() (y int) {
	x := 3
	defer func() {
		//x虽然 + 1 = 4，单并不是返回值变量y
		x++
		fmt.Println("2、defer f3() x=", x)
	}()
	fmt.Println("1、f3() x=", x)
	return x
	//把变量x = 3 赋值给y = 3，因此返回值是3.
}

func f4() (x int) {
	fmt.Println("1. 初始 x:", x) // 0
	//1.函数开始执行，命名返回值 x 初始化为 0
	//为什么返回值是4？因为 defer 匿名函数接收的是 x 的值拷贝，不是原始 x 的引用
	//2.遇到 defer 语句，立即计算并保存 defer 函数的参数值（此时 x 是 0）
	defer func(x int) {
		fmt.Println("3. defer 内的 x:", x) // 0
		x++
		fmt.Println("4. defer 内修改后的 x:", x) // 1
	}(x) //也就是这个 (x) 拷贝了0
	x = 4
	fmt.Println("2. return 前的 x:", x) // 4
	return x
}

/*
内置函数介绍
内置函数			介绍
close			主要用来关闭channel
len				用来求长度，比如string、array、slice、map、channel
new				用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
make			用来分配内存，主要用来分配引用类型，比如chan、map、slice
append			用来追加元素到数组、slice中
panic和recover	用来做错误处理（不推荐用于常规错误处理）
*/
func funcA() {
	fmt.Println("func A")
}

/*
程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。
这个时候我们就可以通过recover将程序恢复回来，继续往后执行。
PS：panic和recover（不推荐用于常规错误处理，对于预期可能发生的错误，使用error返回值）
*/
func funcB() {
	defer func() {
		//recover则类似 java 的 try...catch 对异常进行捕获并进行后续处理。因此 panic 和 recover 搭配使用保证某些情况，程序即使报错，也不会让后续代码无法执行。
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	//类似 java 的 throw 抛出异常
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

// 对于预期可能发生的错误，使用error返回值
func funcD(b bool) (result string, err error) {
	if b {
		return "", errors.New("出错误啦~")
	}
	return "正常返回", nil
}
