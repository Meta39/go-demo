package main

import "fmt"

/*
ProcessControl 流程控制
流程控制是每种编程语言控制逻辑走向和执行次序的重要部分，流程控制可以说是一门语言的“经脉”。
Go语言中最常用的流程控制有if和for，而switch和goto主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。
*/
func ProcessControl() {
	fmt.Println("============ 流程控制 ============")

	/*
		if else(分支结构)
		if条件判断基本写法
		Go语言中if条件判断的格式如下：
		if 表达式1 {
			分支1
		} else if 表达式2 {
			分支2
		} else{
			分支3
		}
		当表达式1的结果为true时，执行分支1，否则判断表达式2，如果满足则执行分支2，都不满足时，则执行分支3。
		if判断中的else if和else都是可选的，可以根据实际需要进行选择。
		Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。
		同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。
		举个例子：
	*/
	ifDemo1()

	/*
		if条件判断特殊写法
		if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，举个例子：
	*/
	ifDemo2()

	/*
		for(循环结构)
		Go 语言中的所有循环类型均可以使用for关键字来完成。
		for循环的基本格式如下：
		for 初始语句;条件表达式;结束语句{
			循环体语句
		}
		条件表达式返回true时循环体不停地进行循环，直到条件表达式返回false时自动退出循环。
	*/
	//常规for的写法
	forDemo()
	//for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	forDemo2()
	//for循环的初始语句和结束语句都可以省略
	forDemo3()
	//for无限循环。可以通过break、goto、return、panic语句强制退出循环。
	forDemo4()

	/*
		for range(键值循环)
		Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
		数组、切片、字符串返回索引和值。
		map返回键和值。
		通道（channel）只返回通道内的值。
		Go1.22版本开始支持 for range 整数。
	*/
	forRangeDemo()
	//for range 整数
	forRangeDemo2()

	/*
		switch case
		使用switch语句可方便地对大量的值进行条件判断。
		case语句中不需要break，java中case需要break，否则后续的语句还会执行，但是go不一样，go不需要break。当然了，也可以写break。
	*/
	switchDemo()
	//Go语言规定每个switch只能有一个default分支。一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switchDemo2()
	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。
	switchDemo3()
	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。一般情况下用不上。
	switchDemo4()

	/*
		goto(跳转到指定标签)
		goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程。
	*/
	//双层嵌套的for循环要退出（可以用下面的goto简化）
	gotoDemo1()
	//使用goto语句能简化代码
	gotoDemo2()

	/*
		break(跳出循环)
		break语句可以结束for、switch和select的代码块。
		break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。
	*/
	breakDemo()

	/*
		continue(继续下次循环)
		continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
		在 continue语句后添加标签时，表示开始标签对应的循环。
	*/
	continueDemo()

	fmt.Println("============ 流程控制 ============")
}

// if条件判断常规写法
func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//可以使用 score
	fmt.Println("变量不在if条件内，if外面可以使用。", score)
}

// if条件判断特殊写法
func ifDemo2() {
	//如果 score 写到if条件里面，那么 score 是无法在if条件外面使用的。
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//无法使用 score
	//fmt.Println(score)
	fmt.Println("变量在if条件内，if外面不可以使用score。因此无法输出结果")
}

// 条件表达式返回true时循环体不停地进行循环，直到条件表达式返回false时自动退出循环。
func forDemo() {
	fmt.Println("常规for的写法")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// for循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：
func forDemo2() {
	fmt.Println("for循环的初始语句可以被忽略，但是初始语句后的分号必须要写的写法")
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
}

/*
forDemo3 for循环的初始语句和结束语句都可以省略，例如：
这种写法类似于其他编程语言中的while，在while后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。
*/
func forDemo3() {
	fmt.Println("for循环的初始语句和结束语句都可以省略的写法")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

/*
无限循环

	for {
	    循环体语句
	}

for循环可以通过break、goto、return、panic语句强制退出循环。
*/
func forDemo4() {
	fmt.Println("for循环可以通过break、goto、return、panic语句强制退出循环。")
	i := 0
	for {
		if i == 10 {
			fmt.Println(i)
			break //跳出循环，return是返回结果，如果for后面还有代码，则都不执行。但是break会执行for后面的代码。
		} else {
			i++
		}
	}
}

// for range(键值循环)常规写法
func forRangeDemo() {
	fmt.Println("Go1.22版本前，for range(键值循环)常规写法。Go1.22版本后，如果不需要i，则推荐for range 5 {//省略内容}。如果需要i还是使用当前写法")
	for i := range 5 {
		fmt.Println(i)
	}
}

// for range(键值循环)常规写法
func forRangeDemo2() {
	fmt.Println("Go1.22版本后推荐当前写法，除非需要i，否则直接使用此写法。类似java的foreach写法")
	for range 2 {
		fmt.Println("当前内容循环输出2次")
	}
}

/*
使用 switch 语句可方便地对大量的值进行条件判断。
1、性能比if...else if...else更好
2、case语句中不需要break，java中case需要break，否则后续的语句还会执行，但是go不一样，go不需要break。当然了，也可以写break。
*/
func switchDemo() {
	fmt.Println("switch case语句中不需要break，java中case需要break，否则后续的语句还会执行，但是go不一样，go不需要break。当然了，也可以写break。")
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
		//break//可以写，可以不写。
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

/*
Go语言规定每个switch只能有一个default分支。
一个分支可以有多个值，多个case值中间使用英文逗号分隔。
*/
func switchDemo2() {
	fmt.Println("Go语言规定每个switch只能有一个default分支。一个分支可以有多个值，多个case值中间使用英文逗号分隔。")
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

/*
分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
*/
func switchDemo3() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

/*
fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。一般情况下用不上。
*/
func switchDemo4() {
	fmt.Println("fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。一般情况下用不上。")
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough //一般情况下用不上。
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

/*
双层嵌套的for循环要退出（可以用下面的goto简化）
*/
func gotoDemo1() {
	fmt.Println("双层嵌套的for循环要退出（可以用下面的goto简化）")
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			fmt.Println("结束for循环")
			break
		}
	}
}

/*
使用goto语句简化代码
*/
func gotoDemo2() {
	fmt.Println("goto语句简化代码")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

// break(跳出循环)
func breakDemo() {
	fmt.Println("break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。")
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

// continue(继续下次循环)
func continueDemo() {
	fmt.Println("continue(继续下次循环)。这个跟java是一样的。除了可以在continue语句后添加标签以外，其他都和java一样")
forloop1:
	for i := 0; i < 3; i++ {
		// forloop2:
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				//即：不输出2-2
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
